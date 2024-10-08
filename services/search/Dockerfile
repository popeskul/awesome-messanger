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
RUN go build -o /app/search-service ./cmd/server/main.go

# Start a new stage from scratch
FROM scratch

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/search-service /app/search-service

# Expose ports for gRPC and HTTP
EXPOSE 50050
EXPOSE 8050

# Command to run the executable
ENTRYPOINT ["/app/search-service"]
