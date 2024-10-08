# Start from the official Go image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o /app/profile-service ./cmd/server/main.go

# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/profile-service /app/profile-service

# Copy the swagger directory
COPY swagger /app/swagger

# Set the working directory
WORKDIR /app

# Expose ports for gRPC and HTTP
EXPOSE 50040
EXPOSE 8040
EXPOSE 8041

# Command to run the executable
ENTRYPOINT ["/app/profile-service"]
