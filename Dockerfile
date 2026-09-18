# Build stage
FROM golang:1.20-alpine AS builder

# Set environment variables
ENV GO111MODULE=on

# Create a directory in the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project into the working directory
COPY . .

# Build the Go app
RUN go build -o main .

# Final stage: minimal image to run the application
FROM alpine:latest

WORKDIR /root/

# Copy the prebuilt binary from the build stage
COPY --from=builder /app/main .

EXPOSE 8001

# Run the Go binary
CMD ["./main"]
