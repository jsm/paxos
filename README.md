# Paxos Coding Challenge

## Installation

Clone this repo to `$GOPATH/src/github.com/jsm/paxos`

```
cd $GOPATH/src/github.com/jsm/paxos
dep ensure
```

If you need to install dep see [https://golang.github.io/dep/docs/installation.html](https://golang.github.io/dep/docs/installation.html)

## Challenge 1

Run locally

```
go run c1/main.go
```

Hitting deployed server

```
curl -X POST -H "Content-Type: application/json" -d '{"message": "foo"}' 54.245.159.148/messages

curl 54.245.159.148/messages/2c26b46b68ffc68ff99b453c1d30413413422d706483bfa0f98a5e886266e7ae
```

### Bottlenecks

This would run into memory bottlenecks as more digests are stored. Also, it doesn't survive an application restart.

To scale this, I would store digests into a key-value lookup optimized db, such as a NoSQL store or a finely-tuned Relational DB.

## Challenge 2

Build

```
go build -o /tmp/c2 c2/main.go
```

Run

```
/tmp/c2 /path/to/price/file 2500
```

### Big O

O(nlogn)

## Challenge 2 (Bonus)

Build

```
go build -o /tmp/c2bonus c2bonus/main.go
```

Run

```
/tmp/c2bonus /path/to/price/file 2500
```

## Challenge 3

Build

```
go build -o /tmp/c3 c3/main.go
```

Run

```
/tmp/c3 X0X
```

### Big O

O(2^n) where n is the number of X's

