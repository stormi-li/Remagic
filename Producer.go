package remagic

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	researd "github.com/stormi-li/Researd"
)

type Producer struct {
	researdClient *researd.Client
	channel       string
	maxRetries    int
	conn          net.Conn
}

func newProducer(researdClient *researd.Client, channel string) *Producer {
	producer := Producer{
		researdClient: researdClient,
		maxRetries:    5,
		channel:       channel,
	}
	go producer.connet()
	return &producer
}

func (producer *Producer) connet() {
	producer.researdClient.Discover(producer.channel, func(addr string) {
		conn, err := net.Dial("tcp", addr)
		if err == nil {
			producer.conn = conn
		}
	})
}

func (producer *Producer) SetMaxRetries(maxRetries int) {
	producer.maxRetries = maxRetries
}

func (producer *Producer) Publish(message []byte) error {
	// 设置重试次数限制，避免无限重试
	count := 0
	for producer.conn == nil {
		time.Sleep(2 * time.Second)
		if count == producer.maxRetries {
			return fmt.Errorf("以达到最大重试次数")
		}
		count++
	}
	retryCount := 0
	fullMessage := []byte(string(message))
	messageLength := uint32(len(fullMessage))

	// 1. 写入消息长度前缀
	lengthBuf := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBuf, messageLength)
	for {
		// 尝试写入字节流

		_, err := producer.conn.Write(append(lengthBuf, fullMessage...))
		if err != nil {
			time.Sleep(2 * time.Second)
		} else {
			return nil
		}
		if retryCount == producer.maxRetries {
			break
		}
		retryCount++
	}
	return fmt.Errorf("以达到最大重试次数")
}
