# Use the official Golang image to create a build artifact.
# https://hub.docker.com/_/golang
FROM golang:1.23.2 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies.
# This allows the container build to reuse cached dependencies.
COPY go.* ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o myapp specifies the output file name, replacing the default.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official Debian slim image for a lean production container.
# https://hub.docker.com/_/debian
FROM debian:buster-slim
RUN set -eux; apt-get update; apt-get install -y --no-install-recommends ca-certificates netbase

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp /myapp

# Run the web service on container startup.
CMD ["/myapp"]