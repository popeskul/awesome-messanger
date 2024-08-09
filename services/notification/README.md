# Microservices Deployment with Kubernetes

This repository contains multiple microservices, each with its own Dockerfile and Kubernetes deployment configurations. 

## Prerequisites

- [Docker](https://www.docker.com/)
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/)

## Common Commands

### Starting Minikube

```sh
minikube start
```

### Building and Deploying Services
Each service has its own Makefile that automates the build and deployment process. Below are the steps to build and deploy each service.

### Notification Service

Build and deploy using the following commands:

```sh
make
# or 
make build
make deploy-blue-green
make deploy-canary
```

### Requests

Send a request to the service using the following command:

```sh
grpcurl -v -plaintext -d '{}' localhost:50053 grpc.health.v1.Health/Chec
grpcurl -v -plaintext -d '{"recipient_id": "user123", "message": "You have a new notification!"}' localhost:50053 message.NotificationService/SendNotification
```