package main

import (
	"github.com/go-redis/redis/v8"
	researd "github.com/stormi-li/Researd"
)

var redisAddr = "118.25.196.166:3934"
var password = "12982397StrongPassw0rd"

func main() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: password,
	})
	client := researd.NewClient(redisClient, "remagic-namespace")
	register1 := client.NewRegister("channel-1", "118.25.196.166:8838")
	register2 := client.NewRegister("channel-1", "118.25.196.166:8889")
	register1.UpdateNodeType(researd.Main)
	register2.UpdateNodeType(researd.Standby)
}
