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
make
# or 
make build
make deploy-blue-green
make deploy-canary
```

### Requests

Send a request to the service using the following command:
    
```sh
grpcurl -v -plaintext -d '{}' localhost:50052 grpc.health.v1.Health/Check
grpcurl -v -plaintext -d '{"chat_id": "chat123"}' localhost:50052 message.MessageService/GetMessages
grpcurl -v -plaintext -d '{"sender_id": "user1", "recipient_id": "user2", "content": "Hello, world!"}' localhost:50052 message.MessageService/SendMessage
```
