.PHONY: deploy-kube .build-image kube-apply-configmap .kube-apply-namespace kube-deploy-blue-green \
        kube-deploy-canary kube-status kube-logs kube-service kube-clean vendor .vendor-reset \
        .vendor-google-protobuf .vendor-googleapis .vendor-protovalidate .vendor-tidy port-forward \
        generate .protoc-generate .bin-deps .tidy .kube-apply-ingress .kube-deploy-ingress-controller \
        lint .proto-lint .buf-generate proto-format wire swagger-generate

IMAGE_NAME := message-service
IMAGE_TAG := latest
DOCKER_IMAGE := $(IMAGE_NAME):$(IMAGE_TAG)

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC := $(shell which protoc)
PROTO_PATH := $(CURDIR)/proto
PKG_PROTO_PATH := $(CURDIR)/pkg

LOCAL_BIN := $(CURDIR)/bin

VENDOR_PROTO_PATH := $(CURDIR)/vendor.protobuf

kube-deploy: .build-image .kube-apply-namespace .kube-deploy-ingress-controller \
			 .kube-apply-ingress .kube-deploy-blue-green

.build-image:
	eval $$(minikube docker-env --shell bash) && docker build -t $(DOCKER_IMAGE) .

.kube-apply-configmap:
	kubectl apply -f deploy/configmap.yaml

.kube-apply-namespace:
	kubectl apply -f deploy/namespace.yaml

.kube-deploy-blue-green: .kube-apply-configmap .kube-apply-namespace
	kubectl apply -f deploy/blue-green/message-service-blue-deployment.yaml
	kubectl apply -f deploy/blue-green/message-service-green-deployment.yaml
	kubectl apply -f deploy/blue-green/message-service-service.yaml
	$(MAKE) kube-logs

.kube-deploy-canary: .kube-apply-configmap .kube-apply-namespace
	kubectl apply -f deploy/canary/message-service-deployment.yaml
	kubectl apply -f deploy/canary/message-service-canary-deployment.yaml
	kubectl apply -f deploy/canary/message-service-service.yaml
	kubectl apply -f deploy/canary/message-service-virtualservice.yaml
	$(MAKE) kube-logs

.kube-apply-ingress:
	kubectl apply -f deploy/ingress/message-service-ingress.yaml

.kube-deploy-ingress-controller:
	kubectl apply -f deploy/ingress/nginx-ingress-controller.yaml
	kubectl apply -f deploy/ingress/nginx-ingress-clusterrole.yaml
	kubectl apply -f deploy/ingress/nginx-ingress-clusterrolebinding.yaml

kube-status:
	kubectl get pods
	kubectl get svc

kube-logs:
	@echo "Checking for any issues with the pods..."
	@PODS=$$(kubectl get pods -o name) ; \
	for POD in $$PODS ; do \
	  STATUS=$$(kubectl get $$POD -o jsonpath='{.status.containerStatuses[0].state.waiting.reason}') ; \
	  if [[ "$$STATUS" == "CrashLoopBackOff" || "$$STATUS" == "ErrImageNeverPull" ]]; then \
	    echo "Logs for $$POD:" ; \
	    kubectl logs $$POD ; \
	  fi \
	done

kube-service:
	minikube service message-service

kube-clean:
	kubectl delete -f deploy/configmap.yaml || true
	kubectl delete -f deploy/blue-green/message-service-blue-deployment.yaml || true
	kubectl delete -f deploy/blue-green/message-service-green-deployment.yaml || true
	kubectl delete -f deploy/blue-green/message-service-service.yaml || true
	kubectl delete -f deploy/canary/message-service-deployment.yaml || true
	kubectl delete -f deploy/canary/message-service-canary-deployment.yaml || true
	kubectl delete -f deploy/canary/message-service-service.yaml || true
	kubectl delete -f deploy/canary/message-service-virtualservice.yaml || true

vendor: .vendor-reset .vendor-googleapis .vendor-google-protobuf .vendor-protovalidate \
		.vendor-protoc-gen-openapiv2 .vendor-tidy

.vendor-reset:
	rm -rf $(VENDOR_PROTO_PATH)
	mkdir -p $(VENDOR_PROTO_PATH)

.vendor-tidy:
	find $(VENDOR_PROTO_PATH) -type f ! -name "*.proto" -delete
	find $(VENDOR_PROTO_PATH) -empty -type d -delete

# Set up the vendor directory with the google/protobuf proto files
.vendor-google-protobuf:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/protocolbuffers/protobuf $(VENDOR_PROTO_PATH)/protobuf &&\
	cd $(VENDOR_PROTO_PATH)/protobuf &&\
	git sparse-checkout set --no-cone src/google/protobuf &&\
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/google
	mv $(VENDOR_PROTO_PATH)/protobuf/src/google/protobuf $(VENDOR_PROTO_PATH)/google
	rm -rf $(VENDOR_PROTO_PATH)/protobuf

# Set up the vendor directory with the bufbuild/protovalidate proto files
.vendor-protovalidate:
	git clone -b main --single-branch --depth=1 --filter=tree:0 \
		https://github.com/bufbuild/protovalidate $(VENDOR_PROTO_PATH)/protovalidate && \
	cd $(VENDOR_PROTO_PATH)/protovalidate
	git checkout
	mv $(VENDOR_PROTO_PATH)/protovalidate/proto/protovalidate/buf $(VENDOR_PROTO_PATH)
	rm -rf $(VENDOR_PROTO_PATH)/protovalidate

# Set up the vendor directory with the googleapis proto files
.vendor-googleapis:
	git clone -b master --single-branch -n --depth=1 --filter=tree:0 \
		https://github.com/googleapis/googleapis $(VENDOR_PROTO_PATH)/googleapis &&\
	cd $(VENDOR_PROTO_PATH)/googleapis &&\
	git checkout
	mv $(VENDOR_PROTO_PATH)/googleapis/google $(VENDOR_PROTO_PATH)
	rm -rf $(VENDOR_PROTO_PATH)/googleapis

# Set up the vendor directory with the grpc-gateway protoc-gen-openapiv2 proto files
.vendor-protoc-gen-openapiv2:
	git clone -b main --single-branch -n --depth=1 --filter=tree:0 \
 		https://github.com/grpc-ecosystem/grpc-gateway $(VENDOR_PROTO_PATH)/grpc-gateway && \
 	cd $(VENDOR_PROTO_PATH)/grpc-gateway && \
	git sparse-checkout set --no-cone protoc-gen-openapiv2/options && \
	git checkout
	mkdir -p $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	mv $(VENDOR_PROTO_PATH)/grpc-gateway/protoc-gen-openapiv2/options $(VENDOR_PROTO_PATH)/protoc-gen-openapiv2
	rm -rf $(VENDOR_PROTO_PATH)/grpc-gateway

# Port forward to the pod for local testing
port-forward:
	@echo "Finding pod for service $(IMAGE_NAME)..."
	@POD_NAME=$$(kubectl get pods -l app=$(IMAGE_NAME) -o jsonpath='{.items[0].metadata.name}'); \
	if [ -z "$$POD_NAME" ]; then \
		echo "No pod found for label app=$(IMAGE_NAME)"; \
		exit 1; \
	fi; \
	echo "Port forwarding from local ports to pod $$POD_NAME"; \
	kubectl port-forward pod/$$POD_NAME 8020:8020 & \
	kubectl port-forward pod/$$POD_NAME 8021:8021 & \
	kubectl port-forward pod/$$POD_NAME 50020:50020 &

generate: .bin-deps .buf-generate proto-format .tidy wire swagger-generate gogenerate

SWAGGER_UI_VERSION := 4.15.5
SWAGGER_UI_URL := https://github.com/swagger-api/swagger-ui/archive/v$(SWAGGER_UI_VERSION).tar.gz

.protoc-generate:
	@echo "Generating proto files and Swagger documentation..."
	@mkdir -p $(PKG_PROTO_PATH)
	@$(PROTOC) -I $(CURDIR)/proto \
	   -I $(VENDOR_PROTO_PATH) \
	   --proto_path=$(CURDIR) \
	   --go_out=$(PKG_PROTO_PATH) --go_opt paths=source_relative \
	   --go-grpc_out=$(PKG_PROTO_PATH) --go-grpc_opt paths=source_relative \
	   --grpc-gateway_out=$(PKG_PROTO_PATH) --grpc-gateway_opt paths=source_relative \
	   --grpc-gateway_opt generate_unbound_methods=true \
	   --openapiv2_out=. --openapiv2_opt logtostderr=true \
	   $(PROTO_PATH)/api/message/message_service.proto $(PROTO_PATH)/api/health/health_service.proto
	@echo "Proto files and Swagger documentation generated successfully."

swagger-generate: .protoc-generate
	@echo "Combining Swagger files..."
	@mkdir -p swagger
	@mkdir -p api
	@./scripts/combine_swagger.sh
	@echo "Downloading and extracting Swagger UI..."
	@curl -sL $(SWAGGER_UI_URL) -o swagger-ui.tar.gz
	@tar -xzf swagger-ui.tar.gz
	@cp -r swagger-ui-$(SWAGGER_UI_VERSION)/dist/* swagger/
	@rm -rf swagger-ui-$(SWAGGER_UI_VERSION) swagger-ui.tar.gz
	@echo "Swagger UI files copied to swagger/"
	@echo "Updating Swagger UI configuration..."
	@awk '{gsub("https://petstore.swagger.io/v2/swagger.json", "swagger.swagger.json")}1' swagger/swagger-initializer.js > temp.js && mv temp.js swagger/swagger-initializer.js
	@echo "Swagger UI configuration updated."
	@rm -rf api

.buf-generate:
	$(info run buf generate...)
	PATH="$(LOCAL_BIN):$(PATH)" $(LOCAL_BIN)/buf generate

proto-format:
	$(info run buf format...)
	$(LOCAL_BIN)/buf format -w

# Set up the local bin directory for installing protoc plugins
.bin-deps: export GOBIN := $(LOCAL_BIN)
.bin-deps:
	$(info Installing binary dependencies...)

	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	go install github.com/bufbuild/buf/cmd/buf@v1.32.2
	go install github.com/yoheimuta/protolint/cmd/protolint@latest

.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy

lint: .proto-lint

.proto-lint:
	$(LOCAL_BIN)/protolint -config_path ./.protolint.yaml ./proto/

wire:
	wire ./internal/di/wire.go

gogenerate:
	go generate ./...

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
