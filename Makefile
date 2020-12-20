## Build all binaries 
build:
	$ go build -o bin/golang-fiber-api main/server.go
.PHONY: build

## Run development server
run:
	$ go run main/server.go
.PHONY: run'

## Start compiled app
start:
	$ sh -c './bin/golang-fiber-api'
.PHONY: start

## Start app and watch file changes
dev:
	$ air -c .air.toml
.PHONY: dev

## Run tests
test:
	$ go test -v ./... -tags=test
.PHONY: test

## Get tests coverage
test-ci:
	$ go test -v ./... -coverprofile=coverage.out -tags=test && make test-coverage-page
.PHONY: test-ci 

## Generate page to see coverage visually
test-coverage-page:
	$ go tool cover -html=coverage.out
.PHONY: test-coverage-page

