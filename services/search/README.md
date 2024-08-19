# Microservices Deployment with Kubernetes

This repository contains multiple microservices, each with its own Dockerfile and Kubernetes deployment configurations. The services include `auth`, `friend`, `message`, `notification`, `profile`, and `search`. This guide provides instructions on how to build, deploy, and manage these services using Minikube and Kubernetes.

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

### Search Service

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
query {
  search(nickname: "example_nickname") {
    id
    nickname
    description
  }
}
```