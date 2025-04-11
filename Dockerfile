FROM golang:1.21-alpine

WORKDIR /app

# Copy the entire source code
COPY . .

# Download all dependencies
RUN go mod download

# Build the application
RUN go build -o main .

# Expose port
EXPOSE 8080

# Command to run the application
CMD ["./main"] 