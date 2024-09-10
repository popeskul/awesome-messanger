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
http://localhost:8001/swagger/index.html

grpcurl -d '{}' -plaintext localhost:50000 health.HealthService/Healthz
curl http://localhost:8000/v1/health
curl http://localhost:8000/v1/healthz

grpcurl -d '{}' -plaintext localhost:50000 health.HealthService/Liveness
curl http://localhost:8000/v1/healthz

grpcurl -d '{}' -plaintext localhost:50000 health.HealthService/Check
curl http://localhost:8000/v1/liveness

grpcurl -d '{}' -plaintext localhost:50000 health.HealthService/Readiness
curl http://localhost:8000/v1/readiness

grpcurl -plaintext -d '{"username": "user123", "password": "password123"}' localhost:50000 auth.AuthService/Login
curl -X POST http://localhost:8000/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user123", "password": "password123"}'

grpcurl -plaintext -d '{"token": "your_token_here"}' localhost:50000 auth.AuthService/Logout
curl -X POST http://localhost:8000/v1/login \
  -H "Content-Type: application/json" \
  -d '{"username": "user123", "password": "password123"}'

grpcurl -plaintext -d '{"username": "newuser", "password": "newpassword"}' localhost:50000 auth.AuthService/Register
curl -X POST http://localhost:8000/v1/register \
  -H "Content-Type: application/json" \
  -d '{"username": "newuser", "password": "newpassword"}'

grpcurl -plaintext -d '{"old_token": "your_old_token_here"}' localhost:50000 auth.AuthService/Refresh
curl -X POST http://localhost:8000/v1/register \
  -H "Content-Type: application/json" \
  -d '{"username": "newuser", "password": "newpassword"}'
```