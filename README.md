# ivmanto.com

The website for IVMANTO — a Data & AI consultancy specializing in Google Cloud
Platform. This is a monorepo containing two independently deployed services.

## Repository layout

| Path        | Service           | Stack                        |
| ----------- | ----------------- | ---------------------------- |
| `/` (root)  | Frontend website  | Nuxt 3 (SSG), TypeScript     |
| `/backend`  | HTTP API backend  | Go (module `ivmanto.com/backend`) |

Each service has its own `Dockerfile` and is deployed as a separate Cloud Run
service via `cloudbuild.yaml`.

## Architecture

```
Browser → Nuxt frontend (Cloud Run)
              └─ /api/* → Go backend (Cloud Run)
                              ├─ /api/booking/*  → Google Calendar (DWD impersonation)
                              ├─ /api/blog/*     → GCS bucket (markdown)
                              ├─ /api/contact    → SMTP
                              ├─ /api/ideas      → Vertex AI (Gemini)
                              └─ /api/articles   → in-memory likes counter
```

## Local development

### Frontend

```sh
npm install
npm run dev        # http://localhost:3000  (proxies /api/* to :8080)
```

### Backend

Requires a `backend/.env` file (copy from `backend/.env.example`):

```sh
cd backend && go run ./cmd/server   # http://localhost:8080
```

For GCP APIs locally:

```sh
gcloud auth application-default login
gcloud auth application-default set-quota-project ivmanto-com-prod
```

## Common commands

```sh
# Frontend
npm run dev        # dev server
npm run generate   # static (SSG) build — what Cloud Build runs
npm run format     # Prettier

# Backend
cd backend
go build ./...
go test ./...
```

## Deployment

Pushing to `main` triggers Cloud Build (`cloudbuild.yaml`), which builds and
pushes both Docker images to Artifact Registry and deploys the frontend and
backend to Cloud Run (`europe-west3`, project `ivmanto-com-prod`). Backend
secrets are injected from Secret Manager.

Working branches follow the `dev-vX.Y.Z` pattern and merge to `main` via pull
request; branch protection requires checks to pass.

## Further documentation

- `CLAUDE.md` — detailed architecture and conventions
- `docs/features/dev-vX.Y.Z/` — per-release working notes and release notes
