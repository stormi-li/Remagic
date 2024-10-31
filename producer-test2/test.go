package main

import (
	"strconv"

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
	producer := client.NewProducer("channel-1")
	for i := 0; i < 500; i++ {
		producer.Publish([]byte("bye world" + strconv.Itoa(i)))
	}
}
