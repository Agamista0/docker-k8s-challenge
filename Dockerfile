FROM golang:1.23

# Install dependencies
RUN apt update && apt install -y git

# Create and set the working directory
WORKDIR /app

# Copy the application code & dependencies
COPY app/ .
# Get Go dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o myapp main.go

# Expose the app port
EXPOSE 80
# Run the Go application
CMD ["/app/myapp"]