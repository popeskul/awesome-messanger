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

### Building Docker Images

Setting Docker Environment to Minikube
Before building the Docker images, set the Docker environment to Minikube:

```sh
eval $(minikube docker-env)
```

### Building and Deploying Services
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
grpcurl -v -plaintext -d '{}' localhost:50051 grpc.health.v1.Health/Check
grpcurl -v -plaintext -d '{"username": "user123", "password": "password123"}' localhost:50051 auth.AuthService/Login
grpcurl -v -plaintext -d '{"token": "your_token_here"}' localhost:50051 auth.AuthService/Logout
grpcurl -v -plaintext -d '{"username": "newuser", "password": "newpassword"}' localhost:50051 auth.AuthService/Register
grpcurl -v -plaintext -d '{"old_token": "your_old_token_here"}' localhost:50051 auth.AuthService/Refresh
```