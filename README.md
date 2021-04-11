![Tests Action Status](https://github.com/lucaswilliameufrasio/golang-fiber-api/workflows/Run%20tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lucaswilliameufrasio/golang-fiber-api/badge.svg?branch=main)](https://coveralls.io/github/lucaswilliameufrasio/golang-fiber-api?branch=main)

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