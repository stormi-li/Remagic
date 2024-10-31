package main

import (
	"fmt"
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
	go func(){
		client := remagic.NewClient(redisClient, "remagic-namespace")
		p := client.NewProducer("channel-1")
		for i := 0; i < 500; i++ {
			err := p.Publish([]byte("hello world" + strconv.Itoa(i)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	go func(){
		client := remagic.NewClient(redisClient, "remagic-namespace")
		p := client.NewProducer("channel-1")
		for i := 0; i < 500; i++ {
			err := p.Publish([]byte("3hello world" + strconv.Itoa(i)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	go func(){
		client := remagic.NewClient(redisClient, "remagic-namespace")
		p := client.NewProducer("channel-1")
		for i := 0; i < 500; i++ {
			err := p.Publish([]byte("2hello world" + strconv.Itoa(i)))
			if err != nil {
				fmt.Println(err)
			}
		}
	}()
	client := remagic.NewClient(redisClient, "remagic-namespace")
	p := client.NewProducer("channel-1")
	for i := 0; i < 500; i++ {
		err := p.Publish([]byte("1hello world" + strconv.Itoa(i)))
		if err != nil {
			fmt.Println(err)
		}
	}
	select{}
}
