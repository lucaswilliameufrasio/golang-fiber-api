![Tests Action Status](https://github.com/lucaswilliameufrasio/golang-fiber-api/workflows/Run%20tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lucaswilliameufrasio/golang-fiber-api/badge.svg?branch=main)](https://coveralls.io/github/lucaswilliameufrasio/golang-fiber-api?branch=main)


# Golang Fiber API


## Performance tests results:

Command used to start the app:

```
make build && make start
```

![Screenshot_20220125_223136](https://user-images.githubusercontent.com/34021576/151088076-e9c586f2-07d7-4b47-a116-e34024e09a47.png)


Command used to start load test:

```
oha -n 1000 http://localhost:7979/api/ && reset && oha -n 1000000 -c 1000 http://localhost:7979/api/
```

![Performance results with oha](https://user-images.githubusercontent.com/34021576/151087913-a414fe49-e7fa-43b0-ad4b-ac213423d551.png)


## To run in production mode:

``` bash
make build
make start
```

## To run in development mode:

``` bash
make dev
```

## To test websocket run:

``` bash
npx wscat -c localhost:7979/api/ws
```

## To install counterfeiter

```bash
GO111MODULE=off go get -u github.com/maxbrunsfeld/counterfeiter
```

## To generate Test Double. example:

```bash
cd src/data/protocols
counterfeiter ./encrypter.go Encrypter
```