# Use the official Golang image as a builder
FROM golang:1.22.4 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest
ENV PATH=$PATH:/go/bin

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

RUN swag init

# Install Delve debugger
#RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Build the Go app
RUN go build -gcflags "all=-N -l" -o main .

# Start a new stage from scratch
FROM debian:bookworm-slim

# Install CA certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main /app/main

# Copy Delve debugger from the builder stage
#COPY --from=builder /go/bin/dlv /go/bin/dlv

# Expose port 8080 to the outside world
EXPOSE 8080 2345

# Command to run the executable
CMD ["/app/main"]
#CMD ["/go/bin/dlv", "exec", "/app/main", "--headless", "--listen=:2345", "--api-version=2", "--accept-multiclient"]
