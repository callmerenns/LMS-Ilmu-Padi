# Stage 1: Build the Go application
FROM golang:1.21.0 AS builder

WORKDIR /app

# Copy go.mod and go.sum
COPY go.mod go.sum ./

# Download dependencies
RUN go mod tidy

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/ilmu-padi

# Stage 2: Create a small image
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/task-api /app/ilmu-padi

# Ensure the binary has execute permissions
RUN chmod +x /app/ilmu-padi

# Set the entrypoint to the binary
ENTRYPOINT ["/app/ilmu-padi"]
