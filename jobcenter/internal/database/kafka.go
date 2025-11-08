package database

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type KafkaConfig struct {
	Addr          string `json:"Addr", optional"`
	WriteCap      int    `json:"WriteCap", optional"`
	ReadCap       int    `json:"ReadCap", optional"`
	ConsumerGroup string `json:ConsumerGroup", optional"`
}

type KafkaData struct {
	Topic string
	Key   []byte
	Data  []byte
}

type KafkaClient struct {
	w         *kafka.Writer
	r         *kafka.Reader
	readChan  chan KafkaData
	writeChan chan KafkaData
	c         KafkaConfig
	closed    bool
	mutex     sync.Mutex
}

func NewKafkaClient(c KafkaConfig) *KafkaClient {
	return &KafkaClient{
		c: c,
	}
}

func (k *KafkaClient) StartWrite() {
	w := kafka.Writer{
		Addr:     kafka.TCP(k.c.Addr),
		Balancer: &kafka.LeastBytes{},
	}
	k.w = &w
	k.writeChan = make(chan KafkaData, k.c.WriteCap)
	go k.sendKafka()
}

func (k *KafkaClient) sendKafka() {
	for data := range k.writeChan {
		messages := []kafka.Message{
			{
				Topic: data.Topic,
				Key:   data.Key,
				Value: data.Data,
			},
		}

		var err error
		const retries = 3
		success := false
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			for range retries {
				err = k.w.WriteMessages(ctx, messages...)
				if err == nil {
					success = true
					break
				}
				// 超时或者leader 不可用 重试
				if errors.Is(err, kafka.LeaderNotAvailable) || errors.Is(err, context.DeadlineExceeded) {
					time.Sleep(time.Millisecond * 250)
					success = false
				}
				success = false
				log.Printf("kafka send write message err %s \n", err.Error())

				if !success {
					k.Send(data)
				}
			}
		}()
	}
}

func (k *KafkaClient) Send(data KafkaData) {
	defer func() {
		if err := recover(); err != nil {
			k.closed = true
		}
	}()

	k.writeChan <- data
	k.closed = false
}

func (k *KafkaClient) Start(data KafkaData) {
	defer func() {
		if err := recover(); err != nil {
			k.closed = true
		}
	}()
	k.writeChan <- data // kefka写入数据通道
	k.closed = false
}

func (k *KafkaClient) Close() {
	if k.w != nil {
		k.w.Close()
		k.mutex.Lock()
		defer k.mutex.Unlock()
		if !k.closed {
			close(k.writeChan)
			k.closed = true
		}
	}
	if k.r != nil {
		k.r.Close()
	}
}

func (k *KafkaClient) StartRead() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{k.c.Addr},
		GroupID:  k.c.ConsumerGroup,
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})

	k.r = r
	k.readChan = make(chan KafkaData, k.c.ReadCap)
	go k.readMsg()
}

func (k *KafkaClient) readMsg() {
	for {
		m, err := k.r.ReadMessage(context.Background())
		if err != nil {
			logx.Error(err)
			continue
		}

		data := KafkaData{
			Key:  m.Key,
			Data: m.Value,
		}
		k.readChan <- data
	}
}

func (k *KafkaClient) Read() KafkaData {
	msg := <-k.readChan
	return msg
}
