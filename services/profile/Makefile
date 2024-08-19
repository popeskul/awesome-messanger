.PHONY: all build deploy-blue-green deploy-canary logs service status clean

IMAGE_NAME := profile-service
IMAGE_TAG := latest
DOCKER_IMAGE := $(IMAGE_NAME):$(IMAGE_TAG)

all: build deploy-blue-green

build:
	eval $$(minikube docker-env --shell bash) && docker build -t $(DOCKER_IMAGE) .

deploy-blue-green:
	kubectl apply -f deploy/blue-green/profile-service-blue-deployment.yaml
	kubectl apply -f deploy/blue-green/profile-service-green-deployment.yaml
	kubectl apply -f deploy/blue-green/profile-service-service.yaml
	$(MAKE) logs

deploy-canary:
	kubectl apply -f deploy/canary/profile-service-deployment.yaml
	kubectl apply -f deploy/canary/profile-service-canary-deployment.yaml
	kubectl apply -f deploy/canary/profile-service-service.yaml
	kubectl apply -f deploy/canary/profile-service-virtualservice.yaml
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
	minikube service profile-service

clean:
	kubectl delete -f deploy/blue-green/profile-service-blue-deployment.yaml || true
	kubectl delete -f deploy/blue-green/profile-service-green-deployment.yaml || true
	kubectl delete -f deploy/blue-green/profile-service-service.yaml || true
	kubectl delete -f deploy/canary/profile-service-deployment.yaml || true
	kubectl delete -f deploy/canary/profile-service-canary-deployment.yaml || true
	kubectl delete -f deploy/canary/profile-service-service.yaml || true
	kubectl delete -f deploy/canary/profile-service-virtualservice.yaml || true
