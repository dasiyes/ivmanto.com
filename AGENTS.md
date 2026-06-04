# Agents working in this repository

This file is the entrypoint for any AI agent (Claude, Cursor, Codex, or other) opening a task in this repo. Humans should read `README.md` first.

## Required reading, in order

1. `README.md` — repo layout & quickstart
2. `CLAUDE.md` — architecture, request flow, deployment
3. `.agents/rules.md` — STRICT rules, file layout, frontend and backend code standards
4. `.agents/workflows/development.md` — dev loop, common commands, debugging tips
5. `.agents/pr-review-contract.md` — what the reviewer will check on your PR

If you skip any of these you will produce a PR that gets blocked. Read them first.

## Branching & commits

- Working branch: `dev-vX.Y.Z`, matching the next version in `package.json`. Never push directly to `main`. Branch protection enforces this.
- Conventional commits, lowercase, scoped. Examples from `git log`:
  - `feat(services): add 3 featured services with keyword indexing`
  - `fix(seo): remove www.ivmanto.com self-links in privacy policy`
  - `chore(release): prep dev-v0.1.3 — release notes, README, version bump`
  - `docs: track CLAUDE.md project guidance for Claude Code`
- One logical change per PR. Don't bundle unrelated refactors. If you spot something off-task, flag it in the PR body — don't fix it silently.

## Pre-PR checklist (include evidence in the PR body)

Frontend changes:
- [ ] `npm run lint` passes
- [ ] `npm run generate` passes
- [ ] Verified affected page(s) in the browser via `npm run dev`

Backend changes:
- [ ] `go build ./...` passes
- [ ] `go test ./...` passes
- [ ] `go run ./cmd/server` boots — no new missing-env-var errors

Both:
- [ ] No secrets in the diff (scan for: passwords, API keys, OAuth secrets, JWT signing keys, GA API secret, SMTP creds)
- [ ] No changes to `cloudbuild.yaml`, auth/DWD code, or Secret Manager references unless the task explicitly calls for them
- [ ] PR description states what changed and why (1–3 sentences)

## Secrets

NEVER commit:
- `backend/.env` (it's gitignored — keep it that way)
- Real values for any `*_PASS`, `*_SECRET`, `*_KEY`, `*_TOKEN` env var
- OAuth client IDs, JWT signing keys, API keys (Gemini, GA API secret, SMTP)

Production secrets come from Secret Manager (project: `ivmanto-com-prod`), injected by Cloud Build into Cloud Run at deploy time. Local dev secrets go in `backend/.env` (gitignored).

If your task needs a new secret:
1. Name the new env var in the PR description.
2. Add a placeholder to `backend/.env.example`.
3. The human owner adds the real value to Secret Manager + Cloud Build trigger before the deploy.

## Do not touch without explicit task scope

If the task description does not name one of these, leave it alone. If it seems unavoidable, stop and ask the owner.

- `cloudbuild.yaml` — deploys both services
- `backend/internal/gcal/` — DWD impersonation / Meet conference setup
- `backend/cmd/server/main.go` startup contract — env-var loading and client initialisation order
- `nuxt.config.ts` runtime config / build hooks
- `plugins/analytics.client.ts` — cookie-consent gating
- The `dist` symlink at repo root (points into `.output/public`)
- Secret Manager values (out of repo)

## Standard workflow

1. **Plan** — write the task to `tasks/todo.md` per `.agents/rules.md`. Wait for owner approval before implementing.
2. **Implement** — keep changes minimal. Follow the file layout in `CLAUDE.md` and the conventions in `.agents/rules.md`.
3. **Verify** — run the pre-PR checklist above. Paste the command output into your PR body.
4. **PR** — open against `main`. Fill in `.github/pull_request_template.md`. Link to the `tasks/todo.md` entry.
5. **Review** — Claude reviews using `.agents/pr-review-contract.md`. Address feedback or push back with reasoning.
6. **Merge** — the human owner merges. Reviewers do not merge.

## When unsure

- Re-read the STRICT section of `.agents/rules.md`.
- Check `CLAUDE.md` for architecture context.
- Browse `docs/` for feature history.
- Ask the owner via a PR comment — don't guess on auth, deployment, or architectural questions.
