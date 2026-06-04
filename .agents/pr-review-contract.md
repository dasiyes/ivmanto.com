# PR review contract

This file describes what the reviewer agent (Claude) checks on every PR opened against `main`. It is the bar an agent author can expect to be held to.

## What gets reviewed

Every PR receives a structured review covering:

1. **STRICT rules** — the 6 STRICT rules in [.agents/rules.md](.agents/rules.md): content preservation, blog system integrity, booking system integrity, cookie-consent gating, no unauthorised changes, API proxy pattern.
2. **Pre-PR checklist** — was lint / test / build / generate evidence included in the PR body? If not, the review blocks until the author re-runs and pastes output.
3. **Secret hygiene** — the diff is scanned for committed secrets (passwords, API keys, OAuth secrets, JWT signing keys, GA API secret, SMTP creds).
4. **Scope** — does the diff stay within the stated task scope, or did the author quietly refactor unrelated files?
5. **Code correctness** — logic, error handling, edge cases on the changed paths.
6. **Code quality** — readability, naming, dead code, duplicated logic. Flagged as comments, not blocks.

## Block bar — review requests changes

The PR is blocked if any of these hold:

- A STRICT rule from `.agents/rules.md` is violated.
- A secret value (real password, API key, OAuth secret, JWT signing key, GA API secret, SMTP creds) is committed.
- `cloudbuild.yaml`, `backend/internal/gcal/` auth code, `backend/cmd/server/main.go` startup contract, Secret Manager references, or cookie-consent gating is touched without the task explicitly calling for it.
- Frontend PR without evidence that `npm run lint` and `npm run generate` pass.
- Backend PR without evidence that `go build ./...` and `go test ./...` pass.
- Commit messages don't follow the conventional-commits style used in `git log` (e.g. `feat(services): ...`, `fix(seo): ...`).
- The branch is not `dev-vX.Y.Z` merging into `main`.
- A correctness bug is identified.

## Comment-only — does not block

- Naming, style, readability suggestions.
- Out-of-scope cleanups (flagged with `[SCOPE]` tag — the author should split into a separate PR).
- Missing tests where the package has an existing test pattern — flagged but not blocked unless the change is high-risk (auth, payments, booking, deploy).
- Documentation suggestions.

## What the reviewer will not do

- Merge the PR — that decision stays with the human owner.
- Push commits directly to the PR branch — feedback comes as PR comments only.
- Approve a PR with a flagged STRICT-rule violation, even if the author asks. The owner can override by merging directly.

## Review output format

For every PR the reviewer leaves:

1. **Summary line** — one of:
   - `✅ Looks good — no blockers` (may still include `[STYLE]` suggestions)
   - `⚠️ Needs changes — N blockers` (lists the blockers up top)
   - `🚧 Out of scope — needs owner triage` (when the PR mixes scopes or the task isn't in `tasks/todo.md`)
2. **Inline comments** on specific lines, tagged with the issue category:
   - `[STRICT]` — STRICT-rule violation (blocking)
   - `[BUG]` — correctness issue (blocking)
   - `[SECRET]` — secret in diff (blocking)
   - `[SCOPE]` — out of scope (comment-only)
   - `[STYLE]` — style/readability (comment-only)
   - `[TEST]` — missing test coverage (usually comment-only)
3. **Pre-PR checklist verification** — confirms which items were evidenced in the PR body and which were missing.

## Re-review on changes

When the PR author pushes new commits to address feedback, the reviewer re-runs the contract against the new diff. A blocking comment is only resolved when the underlying issue is fixed — not when the author replies "done".
