package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	remagic "github.com/stormi-li/Remagic"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := remagic.NewClient(redisClient, "remagic-namespace")
	consumer := client.NewConsumer("channel-1", "118.25.196.166:8889", 3)
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
