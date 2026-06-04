---
description: Workflow Orchestration
---

1. Plan Node Default
 * Enter plan mode for ANY non-trivial task (3+ steps or architectural decisions)
 * If something goes sideways, STOP and re-plan immediately - don't keep pushing
 * Use plan mode for verification steps, not just building
 * Write detailed specs upfront to reduce ambiguity

2. Subagent Strategy
 * Use subagents liberally to keep main context window clean
 * Offload research, exploration, and parallel analysis to subagents
 * For complex problems, throw more compute at it via subagents
 * One task per subagent for focused execution

3. Self-Improvement Loop
 * After ANY correction from the user: update tasks/lessons.md with the pattern
 * Write rules for yourself that prevent the same mistake
 * Ruthlessly iterate on these lessons until the mistake rate drops
 * Review lessons at session start for the relevant project

4. Verification Before Done
 * Never mark a task complete without proving it works
 * Diff behaviour between main and your changes when relevant
 * Ask yourself: "Would a staff engineer approve this?"
 * Run tests, check logs, demonstrate correctness

5. Demand Elegance (Balanced)
 * For non-trivial changes: pause and ask, "Is there a more elegant way?"
 * If a fix feels hacky: "Knowing everything I know now, implement the elegant solution"
 * Skip this for simple, obvious fixes - don't over-engineer
 * Challenge your own work before presenting it

6. Autonomous Bug Fixing
 * When given a bug report: just fix it. Don't ask for hand-holding
 * Point at logs, errors, failing tests - then resolve them
 * Zero context switching required from the user
 * Go fix failing CI tests without being told how


## Developer Workflows

### Common Commands

Frontend:
- `npm run dev` — dev server on :3000, proxies `/api/*` to :8080
- `npm run lint` — ESLint with auto-fix
- `npm run format` — Prettier
- `npm run generate` — SSG build (what Cloud Build runs)
- `npm run preview` — preview the generated site

Backend:
- `cd backend && go run ./cmd/server` — start API on :8080
- `go build ./...` — compile all packages
- `go test ./...` — run all tests
- `go test ./internal/gcal/...` — run a single package's tests
- `go vet ./...` — static analysis

GCP (local dev):
- `gcloud auth application-default login` — set up ADC for the local Go binary
- `gcloud auth application-default set-quota-project ivmanto-com-prod` — set the project ADC will bill

### Typical dev loop

1. Pull `main`, branch `dev-vX.Y.Z` matching the next version in `package.json`.
2. Add your plan to `tasks/todo.md`. Wait for owner approval before implementing.
3. Implement minimally per `.agents/rules.md`.
4. Run the relevant verification commands above.
5. Commit using conventional-commits style (see `git log --oneline` for examples).
6. Open a PR against `main` with `.github/pull_request_template.md` filled in.
7. Address review feedback. Re-run verification on any new commits.

### Debugging tips

- Backend won't start? Check the log for `"missing required environment variables"` — that's `internal/config` telling you to update `backend/.env`.
- Frontend `/api/*` calls fail in dev? Confirm the backend is running on :8080 — the Nuxt dev proxy expects it.
- Calendar API 401 / 403? DWD impersonation needs the SA's OAuth client ID authorised in Workspace Admin Console with scope `https://www.googleapis.com/auth/calendar`. Don't change the auth strategy to "fix" this.
- Cloud Run deploys but blog/booking returns 5xx? Check Secret Manager values are present and the Cloud Build trigger references them.
