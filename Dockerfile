# Stage 1: Build the Go application
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code and protobuf files from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Stage 2: Create a small image with the compiled Go binary
FROM alpine:3.18

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Expose the necessary ports (50051 for gRPC and 2112 for Prometheus metrics)
EXPOSE 50051
EXPOSE 2112

# Command to run the executable
CMD ["./main"]
