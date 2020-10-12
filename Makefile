## Build all binaries 
build:
	$ go build -o bin/golang-fiber-api main/main.go
.PHONY: build

## Run development server
run:
	$ go run main/main.go
.PHONY: run'

## Start compiled app
start:
	$ sh -c './bin/golang-fiber-api'
.PHONY: start

## Run tests
test:
	$ go test -v ./... -tags=test
.PHONY: test 