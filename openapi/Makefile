.PHONY: backend-deploy backend-build backend-deploy backend-promote reserve-host-name esp-enable esp-deploy esp-run-build-image esp-run-deploy

ESP_SERVICE_NAME     = my-cloud-endpoints
BACKEND_SERVICE_NAME = my-cloud-endpoints-backend

backend: backend.go
	GOOS=linux GOARCH=amd64 go build -o backend backend.go

backend-build: backend
	cd deploy && \
	skaffold build \
  --filename skaffold.yaml \
  --default-repo "$(REGISTORY)/my-cloud-endpoints/backend" \
  --file-output artifacts.json

backend-deploy:
	cd deploy && \
    gcloud deploy releases create ${APP_VERSION_HYPHENS} \
    --region=asia-northeast1 \
    --delivery-pipeline my-cloud-deploy-pipeline \
    --gcs-source-staging-dir "gs://${ESP_PROJECT_ID}/my-cloud-deploy/source" \
    --build-artifacts artifacts.json \
    --skaffold-file skaffold.yaml \
    --enable-initial-rollout

backend-promote:
	gcloud deploy releases promote \
    --release ${APP_VERSION_HYPHENS} \
    --delivery-pipeline my-cloud-deploy-pipeline \
    --region=asia-northeast1

esp-reserve-host-name:
	gcloud run deploy $(ESP_SERVICE_NAME) \
    --image="gcr.io/cloudrun/hello" \
    --allow-unauthenticated \
    --platform managed \
    --region=asia-northeast1 \
    --project=$(ESP_PROJECT_ID)

esp-enable:
	gcloud services enable $(FULL_ESP_SERVICE_NAME) \

esp-deploy:
	gcloud endpoints services deploy openapi-run.yaml \
    --project $(ESP_PROJECT_ID)

esp-run-build-image:
	./gcloud_build_image.sh -s $(FULL_ESP_SERVICE_NAME) -c $(CONFIG_ID) -p $(ESP_PROJECT_ID)

esp-run-deploy:
	gcloud run deploy $(ESP_SERVICE_NAME) \
  --image="gcr.io/$(ESP_PROJECT_ID)/endpoints-runtime-serverless:$(ESP_VERSION)-$(FULL_ESP_SERVICE_NAME)-$(CONFIG_ID)" \
  --allow-unauthenticated \
  --platform managed \
  --region=asia-northeast1 \
  --project=$(ESP_PROJECT_ID)
