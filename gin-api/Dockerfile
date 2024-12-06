# Stage 1: Build the Go application
FROM golang:1.21 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to cache dependencies
COPY go.mod go.sum ./

# Download and cache Go modules
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application for production
RUN go build -o gin-api .

# Stage 2: Create a lightweight image to run the app
FROM alpine:latest

# Install necessary runtime dependencies
RUN apk --no-cache add ca-certificates

# Set the working directory
WORKDIR /app

# Copy the built application from the builder stage
COPY --from=builder /app/gin-api .

# Expose the port on which the application runs
EXPOSE 8080

# Command to run the application
CMD ["./gin-api"]