# RESEARD Guides

Simple and efficient distributed message queue library.

# Overview

- Only Redis
- Deploy locally
- Every feature comes with tests
- Developer Friendly

# Install

```shell
go get -u github.com/stormi-li/Remagic
```

# Quick Start

```go
package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	remagic "github.com/stormi-li/Remagic"
)

var redisAddr = “localhost:6379”
var password = “your password”

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := remagic.NewClient(redisClient, "remagic-namespace")
	consumer := client.NewConsumer("channel-1", 3, "118.25.196.166:8889")
	go consumer.HandleMessage(func(message []byte) {
		fmt.Println(string(message))
		time.Sleep(2 * time.Millisecond)
	})
	producer := client.NewProducer("channel-1")
	for i := 0; i < 500; i++ {
		producer.Publish([]byte("hellow world" + strconv.Itoa(i)))
	}
	select {}
}
```

# Interface - remagic

## NewClient

### Create remagic client
```go
package main

import (
	"github.com/go-redis/redis/v8"
	remagic "github.com/stormi-li/Remagic"
)

var redisAddr = “localhost:6379”
var password = “your password”

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := remagic.NewClient(redisClient, "remagic-namespace")
}
```
The first parameter is a redis client of successful connection, the second parameter is a unique namespace.

# Interface - remagic.Client

## NewProducer

### Create a new producer
```go
producer := client.NewProducer("channel-1")
```
The parameter is channel name.

## NewConsumer

### Create a new consumer
```go
consumer := client.NewConsumer("channel-1", 3, "localhost:8899")
```
The first parameter is channel name,  the second parameter is host weight, the third parameter is host address.

# Interface - remagic.Producer

## SetMaxRetries

### Set max retries
```go
producer.SetMaxRetries(10)
```
The parameter is max retries, the default value is 5. The total waiting time is maxRetries * 500 * time.Millisecond

## Publish

### Publish message
 ```go
err := producer.Publish([]byte("hellow world"))
```
The parameter is message. When the maximum number of retries is reached, an error is returned. 

# Interface - remagic.Consumer

## SetCapacity

### Set queue chan capacity
```go
consumer.SetCapacity(5000)
```
The parameter is queue chan capacity, default value is 1000. When the number of messages received exceeds the capacity, the excess messages are saved to the buffer slice.

## HandleMessage

### Handle received message
```go
consumer.HandleMessage(func(message []byte) {
  fmt.Println(string(message))
})
```
The parameter is a hanlder for received message.

#  Community

## Ask

### How do I ask a good question?
- Email - 2785782829@qq.com
- Github Issues - https://github.com/stormi-li/Researd/issues