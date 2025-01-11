# Use the official Golang image as a base
FROM golang:1.23.4

# Set environment variables for Go
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Set the working directory
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o main .


# Expose the port
EXPOSE 3000

# Command to run the application
CMD ["./main"]
