![Tests Action Status](https://github.com/lucaswilliameufrasio/golang-fiber-api/workflows/Run%20tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lucaswilliameufrasio/golang-fiber-api/badge.svg?branch=main)](https://coveralls.io/github/lucaswilliameufrasio/golang-fiber-api?branch=main)


# Golang Fiber API


## Performance tests results:


Command used:

```
oha -n 1000 http://localhost:7979/api/ && reset && oha -n 1000000 -c 1000 http://localhost:7979/api/
```


Summary:
  Success rate:	1.0000
  Total:	7.5241 secs
  Slowest:	0.1579 secs
  Fastest:	0.0001 secs
  Average:	0.0074 secs
  Requests/sec:	132905.9865

  Total data:	27.66 MiB
  Size/request:	29 B
  Size/sec:	3.68 MiB

Response time histogram:
  0.002 [138451] |■■■■■■■■■■■■■■■■■■■■
  0.004 [215759] |■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.006 [185169] |■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.008 [136249] |■■■■■■■■■■■■■■■■■■■■
  0.010 [94194]  |■■■■■■■■■■■■■
  0.012 [65899]  |■■■■■■■■■
  0.014 [48384]  |■■■■■■■
  0.016 [33828]  |■■■■■
  0.018 [23192]  |■■■
  0.020 [16315]  |■■
  0.022 [42560]  |■■■■■■

Latency distribution:
  10% in 0.0017 secs
  25% in 0.0032 secs
  50% in 0.0056 secs
  75% in 0.0097 secs
  90% in 0.0150 secs
  95% in 0.0192 secs
  99% in 0.0298 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0429 secs, 0.0002 secs, 0.1459 secs
  DNS-lookup:	0.0000 secs, 0.0000 secs, 0.0002 secs

Status code distribution:
  [200] 1000000 responses

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