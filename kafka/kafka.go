package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

// 专门往kafka中写入日志文件

var (
	client sarama.SyncProducer
)

func Init(addrs []string) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	var err error
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return err
	}
	return nil
}

func SendToKafka(topic, data string) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(data),
	}

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("pid:%v,offset:%v\n", pid, offset)
	return nil
}
