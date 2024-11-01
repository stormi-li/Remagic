package remagic

import (
	"github.com/go-redis/redis/v8"
	"github.com/stormi-li/Remagic/Researd"
)

type Client struct {
	researdClient *researd.Client
}

func NewClient(redisClient *redis.Client, namespace string) *Client {
	return &Client{
		researdClient: researd.NewClient(redisClient, namespace),
	}
}

func (c *Client) NewProducer(channel string) *Producer {
	return newProducer(c.researdClient, channel)
}

func (c *Client) NewConsumer(channel string, weight int, address string) *Consumer {
	return newConsumer(c.researdClient, channel, weight, address)
}