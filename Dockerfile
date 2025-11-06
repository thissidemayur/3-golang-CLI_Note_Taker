######  Stage 1: Build the Go binary
FROM golang:1.24.9-alpine As Builder

# set working directory
WORKDIR /app

# Copy Go modules first so docker can cache dependencies
COPY go.mod go.sum ./

# Download Go dependencies
RUN go mod download

# Copy the rest of code
COPY . . 

# Build the binary for linux
RUN go build -o bin/app ./cmd/note-taker/main.go


##### Stage 2: Run the binary in a lightweight image
FROM alpine:latest
# installation certificates
RUN apk add --no-cache ca-certificates
# set working directory
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=Builder /app/bin/app .

# Run the binary
ENTRYPOINT ["./app"]