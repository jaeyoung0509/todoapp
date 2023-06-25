FROM golang:1.19-alpine as builder 

#Set the working directory
WORKDIR /app

# Copy the go.mod and go.sum files
COPY go.mod .
COPY go.sum .

# Download the dependencies
RUN go mod download

# Copy the rest of source code
COPY . .

# Build the Go app
RUN CGO_ENABLED=0  GOOS=linux go build -o main ./cmd

# Start a new stage to reduce iamge size 
FROM alpine:latest

# Install ca-certificates
RUN apk --no-cache add ca-certificates

# Copy the binary from the previous stage
COPY --from=builder /app/main /main

# Expose the port for the app to run on
EXPOSE 3000

# Run the binary 
CMD ["/main"]