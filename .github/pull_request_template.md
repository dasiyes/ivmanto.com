## Summary

<!-- 1–3 sentences: what changed and why. Focus on the why; the diff shows the what. -->

## Type & scope

<!-- One conventional-commits type and one scope, matching git log:
     type:  feat | fix | chore | docs | refactor | test | perf
     scope: services | seo | backend | blog | booking | frontend | release | ...  -->

- Type:
- Scope:

## Linked task

<!-- Path to the entry in tasks/todo.md, or issue/PR link. -->

## Pre-PR checklist

**Frontend changes** (delete this block if not applicable):
- [ ] `npm run lint` passes — paste last 5 lines of output:
  ```
  ```
- [ ] `npm run generate` passes — paste last 5 lines of output:
  ```
  ```
- [ ] Verified affected page(s) in the browser via `npm run dev`

**Backend changes** (delete this block if not applicable):
- [ ] `go build ./...` passes
- [ ] `go test ./...` passes — paste last 5 lines of output:
  ```
  ```
- [ ] `go run ./cmd/server` boots locally with no new missing-env-var errors

**Always:**
- [ ] No secrets in the diff (`*_PASS`, `*_SECRET`, `*_KEY`, `*_TOKEN`, OAuth, JWT)
- [ ] No unintended changes to `cloudbuild.yaml`, `backend/internal/gcal/`, `backend/cmd/server/main.go` startup, Secret Manager refs, or cookie-consent gating
- [ ] Commit messages follow conventional-commits style (`type(scope): summary`)
- [ ] Branch is `dev-vX.Y.Z`

## New env vars / secrets

<!-- If this PR adds env vars, name them here. Confirm each:
     - placeholder in backend/.env.example
     - destination: cloudbuild.yaml substitution OR Secret Manager
     The owner provisions Secret Manager values before the deploy. -->

None.

## Notes / risks for the reviewer

<!-- Anything the reviewer should pay extra attention to: tricky logic, areas you weren't sure about, follow-ups deferred to another PR. -->
