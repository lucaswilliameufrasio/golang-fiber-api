FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/go-fiber-api/app/
COPY . .
RUN go get -d -v ./...
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/golang-fiber-api src/src/main/server.go

FROM alpine:3.12
COPY --from=builder /go/bin/golang-fiber-api /bin/golang-fiber-api
WORKDIR /bin
CMD ["/bin/golang-fiber-api"]