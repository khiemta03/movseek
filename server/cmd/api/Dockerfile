# Build stage
FROM golang:1.22.7-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the entire project into the container
COPY . .

# Build the application
RUN go build -o api ./cmd/api

# Final stage
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/api .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./api"]
