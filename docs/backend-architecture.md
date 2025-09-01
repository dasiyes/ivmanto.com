# Backend Architecture & Integration Plan

This document outlines the architecture for the Go-based backend service, its integration with the frontend, and the deployment strategy on Google Cloud Platform.

## 1. Overview

The backend will be a Go application responsible for handling tasks like the contact form submissions. It will be deployed as a serverless container on **Cloud Run** and integrated into our existing GCP infrastructure.

- **Language:** Go
- **Deployment:** Docker container on Cloud Run
- **Integration:** Serves as a backend service for the existing Global External HTTPS Load Balancer.
- **Repository:** Lives in the `/backend` directory of the main project repository.

## 2. Monorepo Project Structure

We will create a new top-level `backend` directory. A standard Go project structure will be used.

```
ivmanto.com/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go      # Main application entry point
│   ├── internal/
│   │   └── contact/
│   │       └── handler.go   # HTTP handlers for the contact form
│   ├── Dockerfile           # To containerize the Go app
│   ├── go.mod
│   └── go.sum
├── docs/
│   ├── backend-architecture.md
│   └── deploy-to-GCP.md
├── src/                     # Vue.js frontend code
└── ... (other frontend files)
```

## 3. Production Architecture: "Internal" Communication

The frontend code (JavaScript) runs in the user's browser, which is external to GCP. It cannot _directly_ access GCP's internal network. However, we can orchestrate a seamless and secure connection via the Load Balancer.

**Traffic Flow:**

1.  A user in their browser submits the contact form.
2.  The Vue.js app makes a `POST` request to `https://ivmanto.com/api/contact`.
3.  The request travels over the internet to our **Global External HTTPS Load Balancer**.
4.  The Load Balancer inspects the URL path. It sees `/api/*` and matches our new routing rule.
5.  The Load Balancer routes the request to the **Cloud Run backend service** over Google's private network.
6.  Your Go application in Cloud Run receives the request, processes it, and sends a response back along the same path.

**Load Balancer Configuration:**

```
Host and Path Rules:
  - Host: ivmanto.com
    Paths:
      - Path: /api/*
        Backend: backend-cloud-run-service  (NEW)
      - Path: /* (Default)
        Backend: frontend-storage-bucket    (EXISTING)
```

This architecture is robust, scalable, and secure.

## 4. Local Development Setup

To mimic the production environment locally, we need the frontend dev server to proxy API requests to a locally running Go backend.

1.  **Run Go Backend:** Start the Go application, which will listen on a port like `http://localhost:8080`.
2.  **Run Vue Frontend:** The Vite dev server runs on `http://localhost:5173`.
3.  **Configure Proxy:** We will update `vite.config.js` to proxy any requests to `/api` from the frontend dev server to the backend server.

**Example `vite.config.js` change:**

```javascript
// vite.config.js
import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'

export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      // Proxy /api requests to our local Go backend
      '/api': {
        target: 'http://localhost:8080',
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, ''),
      },
    },
  },
})
```

This allows you to write `fetch('/api/contact')` in your Vue code, and it will work seamlessly in both local development and production.

## 5. API Endpoints

The Go backend will expose the following RESTful endpoints under the `/api` path.

### Contact API

- **`POST /api/contact`**
  - **Description:** Handles submissions from the contact form.
  - **Payload:** `{ "name": string, "email": string, "message": string }`
  - **Response:** `200 OK` on success.

### Booking API (NEW)

- **`GET /api/booking/availability`**
  - **Description:** Fetches available time slots for a given day.
  - **Query Parameters:** `date` (string, format: `YYYY-MM-DD`)
  - **Response:** `200 OK` with a JSON array of time slots: `[{ "startTime": "...", "endTime": "..." }]`

- **`POST /api/booking/book`**
  - **Description:** Creates a new booking for a selected time slot.
  - **Payload:** `{ "startTime": string, "name": string, "email": string, "notes": string }`
  - **Response:** `201-Created` on success with the created booking object. `409 Conflict` if the slot is already taken.

- **`POST /api/booking/cancel`**
  - **Description:** Cancels an existing booking using a cancellation token.
  - **Payload:** `{ "token": string }`
  - **Response:** `200 OK` on success with a confirmation message. `404 Not Found` if the token is invalid or the booking does not exist.

## 6. CI/CD Pipeline with Cloud Build

We will extend our `cloudbuild.yaml` to handle both frontend and backend deployments in a single run.

**New Cloud Build Steps:**

1.  **Build Go Binary:** Compile the Go application.
2.  **Build Docker Image:** Use the `Dockerfile` in the `/backend` directory to build a container image.
3.  **Push to Artifact Registry:** Push the newly built image to Google Artifact Registry.
4.  **Deploy to Cloud Run:** Deploy a new revision of the Cloud Run service using the new image.
5.  **Build Vue App:** (Existing Step) Run `npm run build`.
6.  **Sync to GCS:** (Existing Step) Sync the `dist` folder to the Cloud Storage bucket.

This creates a fully automated, push-to-deploy workflow for the entire application.
