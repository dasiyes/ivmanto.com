Of course. Migrating your website to a new, dedicated Google Cloud Project under your official ivmanto.com account is a critical step for security, billing clarity, and professional organization. I can certainly provide a detailed, step-by-step guide to make this process as smooth as possible.

This guide assumes you are starting with a brand new GCP project and will walk you through setting up the necessary infrastructure, configuring permissions, and deploying your application using the existing cloudbuild.yaml.

Migration Guide: ivmanto.com to a New GCP Project
Phase 1: Google Cloud Project Setup
This phase prepares your new GCP environment.

Step 1: Create a New GCP Project

Go to the Google Cloud Console.
In the top bar, click the project selector dropdown and then click "NEW PROJECT".
Project Name: ivmanto-com-prod (or a name of your choice).
Organization: Select your ivmanto.com organization.
Click "CREATE".
Step 2: Enable Required APIs For your new project, you need to enable the APIs for the services we'll be using.

Navigate to the APIs & Services Dashboard.
Click "+ ENABLE APIS AND SERVICES".
Search for and enable each of the following APIs one by one:
Cloud Build API (for CI/CD)
Cloud Run Admin API (to deploy your services)
Cloud Logging API (to write and view build logs)
Artifact Registry API (to store your Docker images)
Secret Manager API (to securely store your SMTP password)
Google Calendar API (for the booking system)
Vertex AI API (for the AI features)
Step 3: Link Billing Account A billing account is required to use these services.

Go to the Billing section.
If your project isn't already linked, select it and link it to your ivmanto.com billing account.
Phase 2: IAM & Service Account Configuration
This phase sets up the necessary permissions for your automated deployment pipeline and application.

Step 1: Grant Cloud Build Required Roles Cloud Build runs deployments on your behalf. It needs permission to do so.

Go to the Cloud Build settings page.
Ensure the "Service Account" status for the listed service account is Enabled.
Grant the following roles to this service account ([PROJECT_NUMBER]@cloudbuild.gserviceaccount.com):
Cloud Run Admin: Allows deploying to Cloud Run.
Service Account User: Allows the service account to act as the Cloud Run runtime service account.
Step 2: Create a Service Account for the Backend This account will be used by your Go backend to interact with the Google Calendar API.

Go to IAM & Admin -> Service Accounts.
Click "+ CREATE SERVICE ACCOUNT".
Service account name: ivmanto-backend-sa
Click "CREATE AND CONTINUE". On the "Grant this service account access to project" step, add the following roles:

- **Vertex AI User** (to allow access to the GenAI API)
  Click "CONTINUE", then "DONE".

Step 3: Share Your Google Calendar
Go to the `ivmanto-backend-sa` service account details page in the GCP Console.
Copy its email address (e.g., `ivmanto-backend-sa@ivmanto-com-prod.iam.gserviceaccount.com`).
Go to your "IVMANTO Consultations" Google Calendar settings.
Under "Share with specific people or groups", add the copied service account email.
Set its permissions to "Make changes to events".

Step 4: Configure Local Development Environment
To run the backend server on your local machine, you must configure Application Default Credentials (ADC). This allows the Go application to use your personal Google Cloud identity for authentication during development.

Open your terminal and run the following command:
`gcloud auth application-default login`

Follow the prompts to log in with your Google account. This only needs to be done once per machine.

Phase 3: Infrastructure Setup
This phase creates the cloud resources your application will use.

Step 1: Create an Artifact Registry Repository This is where your Docker images will be stored.

Go to Artifact Registry.
Click "CREATE REPOSITORY".
Name: ivmanto-com-repo (this must match \_REPO_NAME in your cloudbuild.yaml).
Format: Select Docker.
Region: Select a region close to you (e.g., europe-west3).
Click "CREATE".
Step 2: Create a Secret in Secret Manager This will securely store your SMTP password.

Go to Secret Manager.
Click "+ CREATE SECRET".
Name: smtp-password (this must match \_SMTP_SECRET_NAME in cloudbuild.yaml).
Secret value: Enter your SMTP password.
Under "Replication", select "User-managed".
Click "ADD REPLICATION" and choose the same region you used for your Artifact Registry (e.g., `europe-west3`). This is a best practice for data locality and ensures your Cloud Run service can access it.
Click "CREATE SECRET".
Step 3: Grant Secret Access to the Backend Service

Go back to the Secret Manager page.
Select the checkbox next to your smtp-password secret.
In the right-hand panel, click "ADD PRINCIPAL".
In the "New principals" field, paste the service account email for your backend service: `ivmanto-backend-sa@<your-project-id>.iam.gserviceaccount.com`.
In the "Select a role" dropdown, choose "Secret Manager Secret Accessor".
Click "SAVE".
Phase 4: Deployment
Now you are ready to deploy.

Step 1: Configure Cloud Build Trigger

Go to the Cloud Build Triggers page.
Connect your GitHub repository to the new GCP project.
Create a new trigger that points to your ivmanto.com repository.
Event: Push to a branch.
Branch: main (or your primary branch).
Configuration: Cloud Build configuration file (cloudbuild.yaml).

> **Critical:** In the "Advanced" or "Service account" section of the trigger configuration, ensure that the service account selected is the **Cloud Build service account** (`[PROJECT_NUMBER]@cloudbuild.gserviceaccount.com`). Do **not** use the default Compute Engine service account. We have already granted the necessary permissions to the correct Cloud Build account in Phase 2. Using the wrong account will cause permission errors.

Substitution Variables: In the "Substitution variables" section of the trigger, you must provide values for the following variables that are specific to your environment. These will override the placeholder defaults in `cloudbuild.yaml`.

- **`_CALENDAR_ID`**: Your "IVMANTO Consultations" calendar ID (e.g., `c_123...abc@group.calendar.google.com`).
- **`_GCAL_AVAILABLE_SLOT_SUMMARY`**: The exact title you use for available slots in your calendar (e.g., `AfB`).
- **`_SMTP_HOST`**: The hostname of your SMTP server (e.g., `smtp.gmail.com`).
- **`_SMTP_PORT`**: The port for your SMTP server (e.g., `587`).
- **`_SEND_FROM`**: The email address the backend will send emails from (e.g., `contact@ivmanto.com`).
- **`_SEND_FROM_ALIAS`**: The "From" name that appears in emails (e.g., `IVMANTO Support`).

The other variables (`_FRONTEND_SERVICE_NAME`, `_BACKEND_SERVICE_NAME`, etc.) have sensible defaults in `cloudbuild.yaml` and do not need to be added here unless you want to change them.

Step 3: Trigger the First Deployment

Make a small commit to your repository (e.g., update a comment in README.md).
Push the commit to your main branch.
Go to the Cloud Build History page and watch the deployment. It will build and push both your frontend and backend images and then deploy them to Cloud Run.
After these steps, your website will be fully migrated and running in the new, dedicated GCP project.
