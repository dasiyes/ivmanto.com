# Auto-Rebuild Frontend: Webhook Trigger Setup

## What was done in code

The Go backend now calls a Cloud Build webhook URL whenever a `.md` article changes in GCS. This avoids the `metadata.json` loop problem from the Pub/Sub approach.

**Files changed:**
- `backend/internal/config/config.go` — added `FrontendRebuildWebhookURL` field
- `backend/internal/blog/handler.go` — added `triggerFrontendRebuild()` method
- `backend/cmd/server/main.go` — passes webhook URL to handler

---

## Setup Steps

### Step 1 — Delete the old Pub/Sub trigger

1. Go to [Cloud Build → Triggers](https://console.cloud.google.com/cloud-build/triggers?project=ivmanto-com-prod)
2. Find `blog-article-rebuild`
3. Click **⋮** → **Delete**

### Step 2 — Create a Webhook trigger

1. Still in **Cloud Build → Triggers**, click **+ CREATE TRIGGER**
2. Fill in:

| Field | Value |
|---|---|
| **Name** | `blog-frontend-rebuild` |
| **Region** | `europe-west3` |
| **Description** | `Triggered by backend when a blog article changes` |
| **Event** | `Webhook event` |

3. Under **Webhook event**:
   - It will auto-create a **Secret** in Secret Manager
   - Note the **Webhook URL preview** — you'll need this in Step 3

4. Under **Source**:

| Field | Value |
|---|---|
| **Repository** | Select your connected GitHub repo |
| **Branch** | `^main$` |

5. Under **Configuration**:

| Field | Value |
|---|---|
| **Type** | `Cloud Build configuration file` |
| **Location** | `Repository` |
| **Config file** | `cloudbuild-frontend-only.yaml` |

6. Click **CREATE**

7. After creation, go back to the trigger details and **copy the full Webhook URL**. It looks like:
   ```
   https://cloudbuild.googleapis.com/v1/projects/ivmanto-com-prod/locations/europe-west3/triggers/blog-frontend-rebuild:webhook?key=AIza...&secret=...
   ```

### Step 3 — Set the env var on the backend Cloud Run service

Run this command (paste your actual webhook URL):

```bash
gcloud run services update ivmanto-backend-service \
  --region=europe-west3 \
  --project=ivmanto-com-prod \
  --update-env-vars="FRONTEND_REBUILD_WEBHOOK_URL=YOUR_WEBHOOK_URL_HERE"
```

> [!IMPORTANT]
> Replace `YOUR_WEBHOOK_URL_HERE` with the full webhook URL from Step 2.
> Also replace `ivmanto-backend-service` with your actual backend service name if different.

### Step 4 — Deploy the backend code changes

Commit and push the Go backend changes to `main`:

```bash
cd /Users/tonevSr/Programming/_proProjects/ivmanto.com
git add backend/
git commit -m "Add Cloud Build webhook trigger for frontend auto-rebuild"
git push origin main
```

This will trigger a full deploy (frontend + backend) via your existing push-to-main trigger.

### Step 5 — Test

After the deploy finishes, re-upload an existing article:

```bash
gsutil cp gs://ivmanto_com_blog_articles/TechnicalDebtAsset.md gs://ivmanto_com_blog_articles/TechnicalDebtAsset.md
```

Then check:
1. [Cloud Build → History](https://console.cloud.google.com/cloud-build/builds?project=ivmanto-com-prod) — should show ONE build from `blog-frontend-rebuild`
2. Backend logs in Cloud Run — should show `"frontend rebuild triggered successfully"`

---

## How it works

```
New .md article uploaded to GCS
        ↓
GCS notification → Pub/Sub → Backend
        ↓
Backend handlePubSubPush:
  1. Filters: is it a .md file? → Yes
  2. Refreshes article cache (existing)
  3. POSTs to Cloud Build webhook URL (new)
        ↓
Cloud Build: frontend-only rebuild (~3 min)
        ↓
New article live on ivmanto.com ✅
```

> [!NOTE]
> `metadata.json` writes to GCS also trigger Pub/Sub, but they go through the same `handlePubSubPush` endpoint which filters them out (not a `.md` file). So no webhook is fired. No loop. ✅
