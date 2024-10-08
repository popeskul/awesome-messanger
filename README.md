## Microservices Deployment with Kubernetes

![Messenger Architecture](./Messenger%20Architecture.jpg)

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

### Building Docker Images

Before building the Docker images, set the Docker environment to Minikube:

```sh
eval $(minikube docker-env)
```

Building and Deploying Services

Each service has its own Makefile that automates the build and deployment process. Below are the steps to build and deploy each service.

### Make Commands

```sh
make compose-up # Run the service locally using Docker Compose
```
```sh
make compose-down # Stop the service running locally
```
```sh
make compose-build # Build the Docker image locally
```
