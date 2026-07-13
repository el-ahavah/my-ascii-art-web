# ---------- Build Stage ----------
FROM golang:1.24-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go module files first
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the application
RUN go build -o ascii-art-web .

# ---------- Runtime Stage ----------
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the executable
COPY --from=builder /app/ascii-art-web .

# Copy required runtime files
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static
COPY --from=builder /app/banners ./banners

# Expose application port
EXPOSE 8080

# Start the application
CMD ["./ascii-art-web"]