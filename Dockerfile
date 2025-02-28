# Use a builder image to compile the Go application
FROM golang:1.23 AS builder

# Install dependencies
RUN apt update && apt install -y git

# Create and set the working directory
WORKDIR /app

# Copy the application code & dependencies
COPY app/ .

# Get Go dependencies
RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp main.go

# Use a smaller base image for the final image
FROM alpine:latest

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp /app/myapp

RUN chmod +x /app/myapp

RUN ls -l /app

# Expose the app port
EXPOSE 80
# Run the Go application
CMD ["/app/myapp"]