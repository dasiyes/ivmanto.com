# Stage 1: Build the Go binary in a dedicated build environment.
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum to leverage Docker layer caching for dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code.
COPY . .

# Build the application into a static binary.
# CGO_ENABLED=0 is crucial for a static binary in Alpine.
# -ldflags="-w -s" strips debug information to create a smaller binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-w -s" -o ivmanto-backend-service ./cmd/server

# Stage 2: Create the final, minimal image using distroless.
FROM gcr.io/distroless/static-debian11

# Copy only the compiled binary from the builder stage.
COPY --from=builder /app/ivmanto-backend-service /

# Set the entrypoint for the container. Cloud Run will provide the PORT.
ENTRYPOINT ["/ivmanto-backend-service"]
