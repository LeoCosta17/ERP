# Build stage
FROM golang:alpine AS builder

WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application
# Disable CGO for a static binary which is safer for Alpine
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/server ./cmd/api

# Final stage
FROM alpine:latest

WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/bin/server /app/server

# Expose the port the app runs on
EXPOSE 9000

# Run the application
CMD ["/app/server"]
