#!/usr/bin/env python3
"""
DWD + Google Meet smoke test.

Mirrors the Go production flow at backend/internal/gcal/calendar.go:
    impersonate.CredentialsTokenSource(ctx, impersonate.CredentialsConfig{
        TargetPrincipal: cfg.GCal.ServiceAccountEmail,
        Scopes:          []string{calendar.CalendarScope},
        Subject:         cfg.GCal.ImpersonateUser,
    })
    ...
    Events.Update(...).ConferenceDataVersion(1).Do()

Auth chain (no service-account key file is used):
    Local ADC  ->  iam.serviceAccounts.signJwt on the SA
                ->  SA acts as the Workspace user via DWD
                ->  Calendar API call signed as that user

Prerequisites:
  1. ADC: `gcloud auth application-default login` (or workload identity).
  2. The ADC principal must have BOTH of these on the target SA:
        roles/iam.serviceAccountTokenCreator
        roles/iam.serviceAccountOpenIdTokenCreator   (sometimes required for signJwt)
     Verify with:
        gcloud iam service-accounts get-iam-policy $GCAL_SA_EMAIL
  3. The SA's OAuth client ID must be authorised in the Workspace Admin Console
     (Security -> Access and data control -> API controls -> Domain-wide delegation)
     for the scope: https://www.googleapis.com/auth/calendar

Install:
    pip install google-auth>=2.18.0 google-api-python-client

Run:
    python dwd_meet_test.py
"""

from __future__ import annotations

import os
import sys
import uuid
from datetime import datetime, timedelta, timezone

import google.auth
from google.auth import impersonated_credentials
from googleapiclient.discovery import build
from googleapiclient.errors import HttpError


# Defaults match the cloudbuild.yaml substitutions used by the Go backend.
SA_EMAIL = os.environ.get(
    "GCAL_SA_EMAIL",
    "ivmanto-backend-sa@ivmanto-com-prod.iam.gserviceaccount.com",
)
SUBJECT_USER = os.environ.get(
    "GCAL_IMPERSONATE_USER",
    "nikolay.tonev@ivmanto.com",
)
CALENDAR_ID = os.environ.get(
    "CALENDAR_ID",
    "c_2950137553d97197f3e7963a9543784e119032ca9cc1b970ea668c6e9d2c9764@group.calendar.google.com",
)
SCOPES = ["https://www.googleapis.com/auth/calendar"]


def build_calendar_service():
    """ADC -> impersonate SA -> act as Workspace user via DWD subject."""
    source_creds, project = google.auth.default(
        scopes=["https://www.googleapis.com/auth/cloud-platform"]
    )
    print(f"[auth] ADC project: {project}", file=sys.stderr)
    print(f"[auth] target SA  : {SA_EMAIL}", file=sys.stderr)
    print(f"[auth] DWD subject: {SUBJECT_USER}", file=sys.stderr)

    target_creds = impersonated_credentials.Credentials(
        source_credentials=source_creds,
        target_principal=SA_EMAIL,
        target_scopes=SCOPES,
        subject=SUBJECT_USER,  # <-- this is the DWD lever (== Go's `Subject`)
    )

    return build("calendar", "v3", credentials=target_creds, cache_discovery=False)


def create_event_with_meet():
    service = build_calendar_service()

    start = datetime.now(timezone.utc) + timedelta(hours=2)
    end = start + timedelta(minutes=30)

    body = {
        "summary": "DWD + Meet sanity check",
        "description": (
            "Created by dwd_meet_test.py. Safe to delete.\n"
            "If this event has a Google Meet link, DWD is working."
        ),
        "start": {"dateTime": start.isoformat()},
        "end": {"dateTime": end.isoformat()},
        # Same shape used in the Go backend (calendar.go:179-187).
        "conferenceData": {
            "createRequest": {
                "requestId": str(uuid.uuid4()),
                "conferenceSolutionKey": {"type": "hangoutsMeet"},
            },
        },
    }

    event = (
        service.events()
        .insert(
            calendarId=CALENDAR_ID,
            body=body,
            conferenceDataVersion=1,  # required, same as Go
            sendUpdates="none",
        )
        .execute()
    )
    return event


def main() -> int:
    try:
        event = create_event_with_meet()
    except HttpError as e:
        print(f"\n[FAIL] HTTP {e.status_code} {e.reason}", file=sys.stderr)
        print(e.content.decode("utf-8", errors="replace"), file=sys.stderr)
        return 1
    except Exception as e:  # noqa: BLE001
        print(f"\n[FAIL] {type(e).__name__}: {e}", file=sys.stderr)
        return 1

    print("\n[OK] DWD + Meet creation succeeded:")
    print(f"  event id   : {event.get('id')}")
    print(f"  html link  : {event.get('htmlLink')}")
    print(f"  meet link  : {event.get('hangoutLink')}")

    cd = event.get("conferenceData") or {}
    entry_points = cd.get("entryPoints") or []
    if entry_points:
        print("  entryPoints:")
        for ep in entry_points:
            print(f"    - {ep.get('entryPointType')}: {ep.get('uri')}")
    return 0


if __name__ == "__main__":
    sys.exit(main())
