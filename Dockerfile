# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/go-fiber-api/app/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v ./...
# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/golang-fiber-api main/server.go
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/golang-fiber-api /go/bin/golang-fiber-api
# Run the hello binary.
ENTRYPOINT ["/go/bin/golang-fiber-api"]