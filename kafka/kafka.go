package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

// 专门往kafka中写入日志文件

type logData struct {
	topic string
	data  string
}

var (
	client      sarama.SyncProducer
	logDataChan chan *logData
)

func Init(addrs []string, maxSize int) error {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	var err error
	client, err = sarama.NewSyncProducer(addrs, config)
	if err != nil {
		return err
	}
	logDataChan = make(chan *logData, maxSize)
	go sendToKafka()
	return nil
}

func SendToChan(topic, data string) {
	msg := &logData{topic: topic, data: data}
	logDataChan <- msg
}

func sendToKafka() {
	for {
		select {
		case ld := <-logDataChan:
			msg := &sarama.ProducerMessage{
				Topic: ld.topic,
				Value: sarama.StringEncoder(ld.data),
			}
			pid, offset, _ := client.SendMessage(msg)
			fmt.Printf("pid:%v,offset:%v\n", pid, offset)
		default:
			time.Sleep(time.Millisecond * 50)
		}
	}
}
