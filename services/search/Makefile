.PHONY: deploy-kube .build-image kube-apply-configmap .kube-apply-namespace kube-deploy-blue-green \
        kube-deploy-canary kube-status kube-logs kube-service kube-clean vendor .vendor-reset \
        .vendor-google-protobuf .vendor-googleapis .vendor-protovalidate .vendor-tidy port-forward \
        generate .protoc-generate .bin-deps .tidy .kube-apply-ingress .kube-deploy-ingress-controller \
        lint .proto-lint .buf-generate proto-format

IMAGE_NAME := search-service
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
	kubectl apply -f deploy/blue-green/search-service-blue-deployment.yaml
	kubectl apply -f deploy/blue-green/search-service-green-deployment.yaml
	kubectl apply -f deploy/blue-green/search-service-service.yaml
	$(MAKE) kube-logs

.kube-deploy-canary: .kube-apply-configmap .kube-apply-namespace
	kubectl apply -f deploy/canary/search-service-deployment.yaml
	kubectl apply -f deploy/canary/search-service-canary-deployment.yaml
	kubectl apply -f deploy/canary/search-service-service.yaml
	kubectl apply -f deploy/canary/search-service-virtualservice.yaml
	$(MAKE) kube-logs

.kube-apply-ingress:
	kubectl apply -f deploy/ingress/search-service-ingress.yaml

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
	minikube service search-service

kube-clean:
	kubectl delete -f deploy/configmap.yaml || true
	kubectl delete -f deploy/blue-green/search-service-blue-deployment.yaml || true
	kubectl delete -f deploy/blue-green/search-service-green-deployment.yaml || true
	kubectl delete -f deploy/blue-green/search-service-service.yaml || true
	kubectl delete -f deploy/canary/search-service-deployment.yaml || true
	kubectl delete -f deploy/canary/search-service-canary-deployment.yaml || true
	kubectl delete -f deploy/canary/search-service-service.yaml || true
	kubectl delete -f deploy/canary/search-service-virtualservice.yaml || true

vendor: .vendor-reset .vendor-tidy

.vendor-reset:
	rm -rf $(VENDOR_PROTO_PATH)
	mkdir -p $(VENDOR_PROTO_PATH)

.vendor-tidy:
	find $(VENDOR_PROTO_PATH) -type f ! -name "*.proto" -delete
	find $(VENDOR_PROTO_PATH) -empty -type d -delete

# Port forward to the pod for local testing
port-forward:
	@echo "Finding pod for service $(IMAGE_NAME)..."
	@POD_NAME=$$(kubectl get pods -l app=$(IMAGE_NAME) -o jsonpath='{.items[0].metadata.name}'); \
	if [ -z "$$POD_NAME" ]; then \
		echo "No pod found for label app=$(IMAGE_NAME)"; \
		exit 1; \
	fi; \
	echo "Port forwarding from local ports to pod $$POD_NAME"; \
	kubectl port-forward pod/$$POD_NAME 8050:8050 &

generate: .tidy

.tidy:
	GOBIN=$(LOCAL_BIN) go mod tidy
