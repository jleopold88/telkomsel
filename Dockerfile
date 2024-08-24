# Use the official Golang image to create a build artifact.
# This image contains Golang and is optimized for building Go applications.
FROM golang:1.22.2 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the Go Modules dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app with static linking
RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM debian:bullseye-slim

# Install ca-certificates
RUN apt-get update && apt-get install -y ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .

# Copy the .env file
COPY .env .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
