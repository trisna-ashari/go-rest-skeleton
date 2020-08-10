# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
LABEL maintainer="Trisna Novi Ashari <trisna.x2@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update
RUN apk upgrade --update-cache --available
RUN apk add git make curl perl bash build-base zlib-dev ucl-dev

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum languages ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
RUN mkdir -p languages
COPY --from=builder /app/main .
COPY --from=builder /app/.env .
COPY --from=builder /app/languages ./languages/

# Expose port 8888 to the outside world
EXPOSE 8888

# Command to run the executable
CMD ["./main"]

# Health check
HEALTHCHECK --interval=60s CMD wget -qO- localhost:8888/ping