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

Build and deploy using the following commands:

```sh
make # run default target
# or 
make build
make deploy-blue-green
make deploy-canary
```

### Requests

Send a request to the service using the following command:
    
```sh
curl -v -X POST "http://localhost:8080/add-friend" \
     -H "Content-Type: application/json" \
     -d '{"userId": "user123", "friendId": "friend456"}'

curl -v -X GET "http://localhost:8080/friends?userId=user123"

curl -v -X GET "http://localhost:8080/live"

curl -v -X GET "http://localhost:8080/ready"

curl -v -X POST "http://localhost:8080/respond-friend-request" \
     -H "Content-Type: application/json" \
     -d '{"friendId": "friend456", "response": "accepted"}'
```