package main

import (
	"fmt"
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
	consumer := client.NewConsumer("channel-1", "118.25.196.166:8888", 3)
	consumer.SetCapacity(50)
	consumer.HandleMessage(func(message []byte) {
		fmt.Println(string(message))
		time.Sleep(10 * time.Millisecond)
	})
}
