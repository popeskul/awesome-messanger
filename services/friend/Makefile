.PHONY: kube-deploy .build-image .kube-apply-configmap .kube-apply-namespace .kube-deploy-blue-green \
		.kube-deploy-canary .kube-status .kube-logs .kube-service kube-clean port-forward generate swagger-generate

IMAGE_NAME := friend-service
IMAGE_TAG := latest
DOCKER_IMAGE := $(IMAGE_NAME):$(IMAGE_TAG)

kube-deploy: .build-image .kube-deploy-blue-green

.build-image:
	eval $$(minikube docker-env --shell bash) && docker build -t $(DOCKER_IMAGE) .

.kube-apply-configmap:
	kubectl apply -f deploy/configmap.yaml

.kube-apply-namespace:
	kubectl apply -f deploy/namespace.yaml

.kube-deploy-blue-green: .kube-apply-configmap .kube-apply-namespace
	kubectl apply -f deploy/blue-green/friend-service-blue-deployment.yaml
	kubectl apply -f deploy/blue-green/friend-service-green-deployment.yaml
	kubectl apply -f deploy/blue-green/friend-service-service.yaml
	$(MAKE) kube-logs

.kube-deploy-canary: .kube-apply-configmap .kube-apply-namespace
	kubectl apply -f deploy/canary/friend-service-deployment.yaml
	kubectl apply -f deploy/canary/friend-service-canary-deployment.yaml
	kubectl apply -f deploy/canary/friend-service-service.yaml
	kubectl apply -f deploy/canary/friend-service-virtualservice.yaml
	$(MAKE) kube-logs

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
	minikube service friend-service

kube-clean:
	kubectl delete -f deploy/configmap.yaml || true
	kubectl delete -f deploy/blue-green/friend-service-blue-deployment.yaml || true
	kubectl delete -f deploy/blue-green/friend-service-green-deployment.yaml || true
	kubectl delete -f deploy/blue-green/friend-service-service.yaml || true
	kubectl delete -f deploy/canary/friend-service-deployment.yaml || true
	kubectl delete -f deploy/canary/friend-service-canary-deployment.yaml || true
	kubectl delete -f deploy/canary/friend-service-service.yaml || true
	kubectl delete -f deploy/canary/friend-service-virtualservice.yaml || true

# Port forward to the pod for local testing
port-forward:
	@echo "Finding pod for service $(IMAGE_NAME)..."
	@POD_NAME=$$(kubectl get pods -l app=$(IMAGE_NAME) -o jsonpath='{.items[0].metadata.name}'); \
	if [ -z "$$POD_NAME" ]; then \
		echo "No pod found for label app=$(IMAGE_NAME)"; \
		exit 1; \
	fi; \
	echo "Port forwarding from local ports to pod $$POD_NAME"; \
	kubectl port-forward pod/$$POD_NAME 8010:8010 8011:8011 &

generate:
	go generate ./...


swagger-generate:
	@swag init -g cmd/server/main.go -o ./swagger

	@mkdir -p swagger-ui

	@curl -L -o swagger-ui/swagger-ui.zip https://github.com/swagger-api/swagger-ui/archive/refs/tags/v4.15.5.zip

	@unzip -o swagger-ui/swagger-ui.zip -d swagger-ui

	@cp -r swagger-ui/swagger-ui-4.15.5/dist/* swagger-ui/

	@rm -rf swagger-ui/swagger-ui-4.15.5 swagger-ui/swagger-ui.zip

	@cp ./swagger/swagger.json ./swagger-ui/swagger.json
	@cp ./swagger/swagger.yaml ./swagger-ui/swagger.yaml

	@rm -rf api

gogenerate:
	go generate ./...

test:
	go test ./...

test-coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

wire:
	wire ./internal/di/wire.go
