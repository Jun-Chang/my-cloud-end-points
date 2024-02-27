.PHONY: backentdpb-gen backend-pb-lint backend-pb-build backend-build backend-deploy esp-reserve-host-name esp-enable esp-deploy esp-run-build-image esp-run-deploy

ESP_SERVICE_NAME     = my-cloud-endpoints
BACKEND_SERVICE_NAME = my-cloud-endpoints-backend

backend-pb-gen:
	buf generate

backend-pb-lint:
	buf lint

api_descriptor.pb:
	buf build -o api_descriptor.pb

backend-pb-build: api_descriptor.pb

backend: backend.go
	GOOS=linux GOARCH=amd64 go build -o backend backend.go

backend-build: backend
	docker build -t $(REGISTORY)/my-cloud-endpoints/grpc/backend .
	docker push $(REGISTORY)/my-cloud-endpoints/grpc/backend

backend-deploy:
	gcloud run deploy $(BACKEND_SERVICE_NAME) \
    --image="$(BACKEND_IMAGE)" \
    --allow-unauthenticated \
    --platform managed \
    --region=asia-northeast1 \
    --use-http2 \
    --project=$(ESP_PROJECT_ID)

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
	gcloud endpoints services deploy api_descriptor.pb grpc-run.yaml \
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
