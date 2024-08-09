.PHONY: all build deploy-blue-green deploy-canary status logs service clean generate-proto

IMAGE_NAME := auth-service
IMAGE_TAG := latest
DOCKER_IMAGE := $(IMAGE_NAME):$(IMAGE_TAG)

PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)
PROTOC := $(shell which protoc)

all: build deploy-blue-green

build:
	eval $$(minikube docker-env --shell bash) && docker build -t $(DOCKER_IMAGE) .

create-configmap:
	kubectl apply -f deploy/configmap.yaml

deploy-blue-green: create-configmap
	kubectl apply -f deploy/blue-green/auth-service-blue-deployment.yaml
	kubectl apply -f deploy/blue-green/auth-service-green-deployment.yaml
	kubectl apply -f deploy/blue-green/auth-service-service.yaml
	$(MAKE) logs

deploy-canary: create-configmap
	kubectl apply -f deploy/canary/auth-service-deployment.yaml
	kubectl apply -f deploy/canary/auth-service-canary-deployment.yaml
	kubectl apply -f deploy/canary/auth-service-service.yaml
	kubectl apply -f deploy/canary/auth-service-virtualservice.yaml
	$(MAKE) logs

status:
	kubectl get pods
	kubectl get svc

logs:
	@echo "Checking for any issues with the pods..."
	@PODS=$$(kubectl get pods -o name) ; \
	for POD in $$PODS ; do \
	  STATUS=$$(kubectl get $$POD -o jsonpath='{.status.containerStatuses[0].state.waiting.reason}') ; \
	  if [[ "$$STATUS" == "CrashLoopBackOff" || "$$STATUS" == "ErrImageNeverPull" ]]; then \
	    echo "Logs for $$POD:" ; \
	    kubectl logs $$POD ; \
	  fi \
	done

service:
	minikube service auth-service

clean:
	kubectl delete -f deploy/configmap.yaml || true
	kubectl delete -f deploy/blue-green/auth-service-blue-deployment.yaml || true
	kubectl delete -f deploy/blue-green/auth-service-green-deployment.yaml || true
	kubectl delete -f deploy/blue-green/auth-service-service.yaml || true
	kubectl delete -f deploy/canary/auth-service-deployment.yaml || true
	kubectl delete -f deploy/canary/auth-service-canary-deployment.yaml || true
	kubectl delete -f deploy/canary/auth-service-service.yaml || true
	kubectl delete -f deploy/canary/auth-service-virtualservice.yaml || true

generate-proto:
	@echo "Generating gRPC code..."
	@mkdir -p pb
	@$(PROTOC) --go_out=pb --go-grpc_out=pb --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/auth_service.proto

# Port forward to the pod for local testing
port-forward:
	@echo "Finding pod for service $(IMAGE_NAME)..."
	@POD_NAME=$$(kubectl get pods -l app=$(IMAGE_NAME) -o jsonpath='{.items[0].metadata.name}'); \
	if [ -z "$$POD_NAME" ]; then \
		echo "No pod found for label app=$(IMAGE_NAME)"; \
		exit 1; \
	fi; \
	echo "Port forwarding from local port 50051 to pod $$POD_NAME"; \
	kubectl port-forward pod/$$POD_NAME 50051:50051 &
