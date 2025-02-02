FROM golang:1.21-alpine

WORKDIR /app

# Install necessary build tools
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY src/go.mod src/go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY src .

# Build the application
RUN go build -o main .

# Run the application
CMD ["/app/main"]