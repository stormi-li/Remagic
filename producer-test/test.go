package main

import (
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
	p := client.NewProducer("channel-1")
	for i := 0; i < 10000; i++ {
		p.Publish([]byte("hello world" + strconv.Itoa(i)))
		time.Sleep(5 * time.Millisecond)
	}
}
