# Stage 1: Build the Go binary
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY ./go.mod ./go.sum ./

# Copy the rest of the application code from the ./class-api directory
COPY ./ ./
COPY ./cmd/api-student/ ./

RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /api-student

# Expose the ports for both REST and gRPC
EXPOSE 8082 50051

# Command to run the binary
CMD ["/api-student"]