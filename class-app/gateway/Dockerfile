# Stage 1: Build the Go binary
FROM golang:1.23.2 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY ./go.mod ./go.sum ./

# Copy the rest of the application code
COPY ./ ./
COPY ../gateway ./

RUN go mod download

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /gateway

# Expose the ports for REST
EXPOSE 8080

# Command to run the binary
CMD ["/gateway"]