# Use an official Go runtime as a parent image for building
FROM golang:1.22-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o /app/friend-service ./cmd/server/main.go

# Use a more feature-rich base image for debugging
FROM alpine:latest

COPY config.yaml /app/config.yaml

# Copy the pre-built binary and swagger files
COPY --from=builder /app/friend-service /app/friend-service
COPY swagger /app/swagger

# Set the working directory
WORKDIR /app

# Expose ports
EXPOSE 8010
EXPOSE 8011

# Command to run the executable
ENTRYPOINT ["/app/friend-service"]
