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
RUN go build -o /app/auth-service ./cmd/server/main.go

# Use a more feature-rich base image for debugging
FROM alpine:latest

# Copy the pre-built binary and swagger files
COPY --from=builder /app/auth-service /app/auth-service
COPY swagger /app/swagger

# Set the working directory
WORKDIR /app

# Expose ports
EXPOSE 50000
EXPOSE 8000
EXPOSE 8001

# Command to run the executable
ENTRYPOINT ["/app/auth-service"]
