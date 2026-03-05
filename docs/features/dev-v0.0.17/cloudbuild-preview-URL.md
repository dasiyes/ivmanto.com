https://cloudbuild.googleapis.com/v1/projects/ivmanto-com-prod/locations/europe-west3/triggers/blog-frontend-rebuild:webhook?key=AIzaSyAuIcqNPvu-BH17AhVPkWKjCB_FCZXZPjg&trigger=blog-frontend-rebuild&projectId=ivmanto-com-prod&secret=<SECRET>

```bash
gcloud run services update ivmanto-backend-service \
  --region=europe-west3 \
  --project=ivmanto-com-prod \
  --update-env-vars="FRONTEND_REBUILD_WEBHOOK_URL=https://cloudbuild.googleapis.com/v1/projects/ivmanto-com-prod/locations/europe-west3/triggers/blog-frontend-rebuild:webhook?key=AIzaSyAuIcqNPvu-BH17AhVPkWKjCB_FCZXZPjg&trigger=blog-frontend-rebuild&projectId=ivmanto-com-prod&secret=<SECRET>"
```
