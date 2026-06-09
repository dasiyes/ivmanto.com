# GCP CI/CD watcher — monitor and control the build/deploy pipe

- **Status:** backlog — pending owner green-light to start
- **Date opened:** 2026-06-09
- **Owner:** Nick (dasiyes)
- **Priority:** deferred — lower than the next Clauco task (awaiting that spec)
- **Related repo:** /home/ivmo/projects/ivmanto.com
- **Affected env:** workstation (this box), project `ivmanto-com-prod` (GCP)

## Why

PR #93 (booking-confirmation TZ) merged and shipped. The local workstation
cannot query GCP from here — no `gcloud` binary installed, only the
config dir at `~/.config/gcloud`. So during the deploy, the only
visibility was: "is `git log` clean on main?" That's not enough for a
production deploy on a customer-facing site. Nick wants the next
iteration to give ivmo end-to-end visibility and control of the
CI/CD pipeline so failures surface to ivmo (and the team via redma)
without the owner having to babysit the Cloud Build console.

## Scope (high level)

1. **gcloud SDK installed on this box** — apt (sudo) or portable tarball
   to `/home/ivmo/google-cloud-sdk/`. No SDK is currently present
   even though `~/.config/gcloud` exists (config-only, no binary).
2. **Service-account auth** — provision an SA in IAM on
   `ivmanto-com-prod` with four roles:
   - `roles/cloudbuild.builds.viewer`
   - `roles/run.developer`
   - `roles/artifactregistry.reader`
   - `roles/logging.viewer`
   Drop the SA key JSON at `/home/ivmo/.config/gcloud/keys/ivmo-cicd.json`
   (chmod 600). Activate with `gcloud auth activate-service-account`.
3. **Four devops skills** added to `~/.hermes/profiles/ivmo/skills/devops/`:
   - `cicd-monitor` — poll `gcloud builds list`, surface status + logs
   - `cicd-rollback` — shift Cloud Run traffic back to the previous revision
   - `cicd-redeploy` — trigger a manual build off a specific commit SHA
   - `cicd-smoke` — post-deploy health check via Cloud Run logs + live URL
4. **(Optional) Cloud Build → Pub/Sub → redma webhook** so a finished
   build auto-broadcasts to `team:general`. Skip in the first pass;
   add if the polling model turns out to miss fast-failing builds.

## Open questions for owner (defer until work starts)

1. **Install method:** apt (faster, needs sudo) or tarball (no sudo)?
2. **SA scope:** four roles as listed, or split into a read-only
   monitor SA + a separate deploy SA?
3. **Skills:** all four at once, or start with `cicd-monitor` only?
4. **Optional webhook (item 4):** in or out of the first pass?

## What's explicitly NOT in scope

- Editing `cloudbuild.yaml` to add a smoke-test gate between image
  push and traffic shift. Worth doing as a separate hardening PR
  (would let `cicd-rollback` be triggered automatically on a failed
  gate, instead of manually). Mentioned by ivmo in the PR #93
  handoff; tracked here for awareness, separate work.
- Touching `cloudbuild.yaml`'s deploy step ordering or DWD-related
  config. Out of scope per `.agents/pr-review-contract.md`'s no-touch
  list.
- Cloud SQL / Secret Manager visibility from the CI pipeline.
  These are intentionally opaque to the build pipeline; SA is not
  scoped for them.

## Verification (when this work starts)

- `gcloud --version` returns a sane version.
- `gcloud builds list --limit 1` returns a real build from
  `ivmanto-com-prod` (proves auth + IAM role work end-to-end).
- `gcloud run services list --region=europe-west3 --project=ivmanto-com-prod`
  lists the frontend and backend services.
- Each of the four skills has a worked example in its SKILL.md
  using the real project name.
- (Manual) Trigger a rollback against a known-good prior revision
  and confirm traffic shifts; shift it back.

## Risks

- **SA key on disk** — single point of compromise. Mitigation: key
  scoped to the four roles above (no admin), rotatable from GCP
  console, chmod 600, never echoed in commands.
- **Service-account scope creep** — `run.developer` can change
  traffic on any service in the project, not just ivmanto.com's
  two services. Acceptable for now; if other workloads land on
  the project, consider project-level separation or per-service
  bindings.
- **Polling latency** — if the optional Pub/Sub webhook is
  skipped, ivmo only sees deploy state when prompted. Acceptable
  for ivmo's working style (the owner tends to issue "watch the
  build" instructions explicitly).
