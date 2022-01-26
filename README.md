![Tests Action Status](https://github.com/lucaswilliameufrasio/golang-fiber-api/workflows/Run%20tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/lucaswilliameufrasio/golang-fiber-api/badge.svg?branch=main)](https://coveralls.io/github/lucaswilliameufrasio/golang-fiber-api?branch=main)


# Golang Fiber API


## Performance tests results (not a real world performance test, just an example):

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


Same test on AWS EKS with two replicas:

![Screenshot_20210611_215838](https://user-images.githubusercontent.com/34021576/151088308-1939b2c2-a3eb-44c0-9525-939874c0604e.png)

At login route:

![Screenshot_20210611_220802](https://user-images.githubusercontent.com/34021576/151088457-40ed7eeb-9c37-42ac-81d2-dc5c9a35ee0b.png)

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