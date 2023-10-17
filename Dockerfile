# Stage 1: Build the application
FROM golang:1.17 AS build

WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of your application source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Stage 2: Create a minimal runtime image
FROM alpine:latest

# Install PostgreSQL client
RUN apk --no-cache add postgresql-client

# Copy the binary built in Stage 1
COPY --from=build /app/main /app/main

# Expose the port that your application listens on
EXPOSE 8080

# Start the application
CMD ["/app/main"]
