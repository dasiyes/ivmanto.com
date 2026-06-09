# Booking confirmation: render time in the visitor's timezone

- **Status:** approved by owner 2026-06-09 — implementation complete, commits `2ce549c` + `9496df2` + `400a00d` pushed to `origin/dev-v0.1.6`. PR #93 open; addressing Clauco review feedback.
- **Date opened:** 2026-06-09
- **Decisions confirmed by owner:**
  - `.ics` rendering is correct (UTC `Z` works as expected in the visitor's calendar). Only the email body needs TZ localisation. The `X-WR-TIMEZONE` hint is a small "while we're here" add — owner explicitly OK'd it.
  - Empty `()` in the email is the same root cause; fix it as part of the same change.
  - Bundle the cancellation-email fix in the same PR.
  - Defaults accepted: Go's `MST` layout for abbreviations, `visitorTimezone` JSON field name.
- **Owner:** Nick (dasiyes)
- **Proposed branch:** `dev-v0.1.6` (current `package.json` is `0.1.3`, but the last merged dev branch is `dev-v0.1.5` per `git log --all --decorate` — sequential +1)
- **Affected services:** backend only (Go)
- **STRICT rules touched:** none — booking system integrity preserved (event still stored in calendar's local time; only the visitor-facing render changes)

## Problem

Confirmed by Nick on 2026-06-09 with a real visitor email:

> Hi Marina Muchakova,
>
> Your 30-minute consultation is confirmed. Here are the details:
>
> Date: Tuesday, June 9, 2026
> Time: 3:30 PM - 4:00 PM ()
> Google Meet Link: https://meet.google.com/pmm-gtyy-ept

The visitor was in a +1 offset (EEST-ish). The wall-clock shown is the **calendar owner's** local time (CEST), not the visitor's. An empty parentheses is also visible — see the "()", which means the timezone string passed to the template was empty (the email's `<strong>Time:</strong> %s - %s (%s)` slot is rendering with an empty `(CEST)` style suffix). That's a secondary, related symptom.

### Why it happens (code path)

- `backend/internal/booking/handler.go:243` — `startTime, _ := time.Parse(time.RFC3339, event.Start.DateTime)` parses the GCal response. `DateTime` includes a `+02:00` offset, so `startTime` is the correct absolute instant.
- `backend/internal/booking/handler.go:258` — `Timezone: startTime.Location().String()`. The location is a fixed `+02:00` zone from RFC3339 parsing, not a named IANA zone. `time.Time.Format("3:04 PM")` then prints `startTime` in *its own location* (the +02:00 fixed zone), giving the owner's wall-clock — not the visitor's. The empty `()` comes from the second effect: the format verb in `smtp.go:213-214` includes `(Timezone)` so the suffix is intended to be the *target* zone, but the handler fills it with the source zone.
- `backend/internal/ical/generator.go:55-56` — `DTSTART`/`DTEND` are emitted as UTC `Z`. **This part is correct** — properly-localised clients will render the event in the user's local timezone. Awaiting Nick's confirmation that the .ics he is checking does, in fact, render at 16:30 EEST in his visitor's calendar.

## Fix scope (minimal, per `.agents/rules.md`)

1. **Frontend (`pages/booking/index.vue`):** include the visitor's IANA timezone in the booking POST body.
   - Source: `Intl.DateTimeFormat().resolvedOptions().timeZone` (already used for the "Timezone: ..." label — split into a separate `timezoneIana` ref so the wire format and the human-readable display don't conflate).
   - Field name: `visitorTimezone` (matches REST naming in `createBookingRequest`).
2. **Backend request struct (`backend/internal/booking/handler.go`):** add `VisitorTimezone string` to `createBookingRequest`.
3. **Backend TZ resolution:** in `handleCreateBooking`, after parsing `event.Start.DateTime`:
   - Attempt `time.LoadLocation(req.VisitorTimezone)`. If it fails (empty string, unknown IANA name, no tzdata in the container), log a warning and fall back to the calendar's location (CEST) — **do not fail the booking**.
   - Build a `time.Time` in the visitor's location (preserve the absolute instant — convert the existing `startTime` to the visitor's zone via `startTime.In(visitorLoc)`, do not re-parse and re-shift).
   - Derive the display label at the call site from the event's actual start time: `startTime.In(visitorLoc).Format("MST")`. This makes the abbreviation DST-correct (EEST in summer, EET in winter for Athens).
4. **Email render (`backend/internal/email/smtp.go`):** pass the resolved visitor location to `details.StartTime`/`EndTime` and the `Timezone` label. Use `time.Time.Format` with the visitor's location already attached (so `Format("3:04 PM")` renders in the visitor's wall-clock).
5. **Cancellation email path:** same fix in the cancellation emails (`SendBookingCancellationToClient`, `SendBookingCancellationToAdmin`) — currently they format `startTime` directly without a TZ context, so a +1 visitor sees organiser time there too. Fix at the same time to avoid a second report.
6. **Visitor TZ persistence on the calendar event:** the cancellation email is triggered by an email link with no live client context, so the visitor IANA must be stored on the calendar event at booking time and read back on cancel. Implemented as `event.ExtendedProperties.Private["visitor_timezone"]` in `gcal.BookingDetails` and surfaced on the `originalEvent` returned from `gcal.CancelBooking`. Privacy: the IANA name itself is not PII (it's roughly equivalent to a coarse geolocation). Owner approved the scope addition in PR review.
7. **iCal (.ics):** add an `X-WR-TIMEZONE:<VisitorIANA>` line at the VCALENDAR level for the visitor's copy. Keep `DTSTART`/`DTEND` as UTC `Z` (already correct). A full `VTIMEZONE` block is out of scope for this PR — `Z` + `X-WR-TIMEZONE` is sufficient for all modern clients.
8. **Tests:** unit tests in `backend/internal/email/` (rendered HTML body for a June Athens scenario) and `backend/internal/booking/handler_test.go` (resolveVisitorTimezone — DST-correct abbreviation, fallback paths, whitespace trimming). Plus the existing `backend/internal/ical/generator_test.go` (X-WR-TIMEZONE presence/absence, escaping).
9. **Defensive logging:** log `visitorTimezone=...` and `resolvedLocation=...` on every booking, so a future bad TZ value is visible in Cloud Run logs.

## Out of scope (flagged for separate PRs)

- Full `VTIMEZONE` block in the .ics (current `Z`-suffixed UTC fields are RFC-5545 compliant and modern clients render correctly).
- Frontend display of the visitor's TZ in the slot grid — the grid is already in the browser's local time, which is the visitor's TZ by definition.

## Scope additions made during implementation (flagged for review)

- **Admin booking notification email (`SendBookingNotificationToAdmin`):** the original plan said "leave as-is, RFC1123 is timezone-explicit". On review, this was wrong: `time.Parse(time.RFC3339, "...+02:00")` returns a time in an unnamed `*time.FixedZone`, and `Format(time.RFC1123)` on that produces an empty/MST-less string. The admin was seeing `Tue, 09 Jun 2026 13:30:00 ` (no suffix). I added a one-line fix in the booking handler to convert the parsed time into the calendar's IANA location before sending to the email service — same root cause, same one-call fix. If you'd rather I drop this, revert the last 8 lines of the `handleCreateBooking` admin goroutine.
- **`backend/internal/gcal/calendar.go`:** the no-touch list in `AGENTS.md` and `.agents/pr-review-contract.md` says "`backend/internal/gcal/` auth code" is off-limits. My edits are to `BookingDetails`, `BookSlot`, and `CancelBooking` (the booking-flow logic), NOT to `NewService` (the DWD auth code). I am calling this out explicitly here and in the PR body so the reviewer can confirm. If the contract is meant to be read more strictly — i.e. the whole `gcal/` package is locked — I will revert these and find a different way (likely persisting the visitor TZ in a separate backend store).

## Pre-PR gates to run

- [ ] `cd backend && go build ./...` — must pass
- [ ] `cd backend && go test ./...` — must pass, including the new test
- [ ] `go run ./cmd/server` boots — must not add new missing-env-var errors
- [ ] Local smoke: with backend running and a real GCal test event, POST to `/api/booking/book` with a `visitorTimezone` of `Europe/Athens` and verify the resulting email body says "4:30 PM - 5:00 PM (EEST)" for an event originally at 15:30 CEST. (Will need a sandbox visitor inbox.)
- [ ] No secrets in the diff
- [ ] No edits to `cloudbuild.yaml`, `backend/internal/gcal/` (only `booking/` and `email/` and `ical/` touched), `backend/cmd/server/main.go`, or `nuxt.config.ts`
- [ ] Conventional commit message: `fix(booking): render confirmation email in visitor's timezone`

## Open questions for owner

1. **Abbreviation source:** `time.Time.Format` with `MST` layout gives zone abbreviations from Go's embedded `tzdata` (since Go 1.15 this is always available, no extra import needed). I will use that. If a non-ASCII/non-standard name is needed, we'll iterate. OK?
2. **Frontend field naming:** I'll go with `visitorTimezone` (camelCase) in the JSON to match the existing `eventId`/`ga_client_id` style. OK?
3. **Empty parentheses in current email:** this is a cosmetic symptom of the same root cause (handler passing the source zone but the format string expecting a display label). Will be fixed by the same change. No separate PR needed unless you want one.

## Review notes

(To be filled in after PR review.)
