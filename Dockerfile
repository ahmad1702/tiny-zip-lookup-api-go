# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

# Stage 2: Run the binary
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary and necessary directories
COPY --from=builder /app/app .
COPY --from=builder /app/data ./data
COPY --from=builder /app/seed ./seed

# Expose the application's port
EXPOSE 8080

# Run the application
CMD ["./app"]
