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

## Update all dependencies
update-deps:
	$ go get -u ./...
.PHONY: update-deps

## Update all test dependencies
update-test-deps:
	$ go get -t -u ./...
.PHONY: update-test-deps

## Build image using root permission
docker-build-sudo:
	$ sudo docker build -t lucaswilliam/go-fiber-api .
.PHONY: docker-build-sudo

## Run image using root permissions
docker-run-sudo:
	$ sudo docker run -d -p 7979:8888 -e PORT=8888 lucaswilliam/go-fiber-api
.PHONY: docker-run-sudo


