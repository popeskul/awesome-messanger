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

### Profile Service

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
grpcurl -v -plaintext -d '{}' localhost:50054 grpc.health.v1.Health/Check
grpcurl -v -plaintext -d '{"user_id": "123", "nickname": "JohnDoe", "bio": "Developer", "avatar_url": "http://example.com/avatar.jpg"}' localhost:50054 message.ProfileService/UpdateProfile
grpcurl -v -plaintext -d '{"user_id": "123"}' localhost:50054 message.ProfileService/GetProfile
```