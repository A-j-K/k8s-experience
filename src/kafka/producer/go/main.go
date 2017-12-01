package main

import (
	"fmt"
	"os"
	"time"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
	"github.com/Shopify/sarama"
)

var (
	brokerList = kingpin.Flag("brokerList", "List of brokers to connect").Default("localhost:9092").Strings()
	topic      = kingpin.Flag("topic", "Topic name").Default("test").String()
)

func main() {
	kingpin.Parse()
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(*brokerList, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			panic(err)
		}
	}()
	for {
		msg := &sarama.ProducerMessage{
			Topic: *topic,
			Value: sarama.StringEncoder(name),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", *topic, partition, offset)
		duration := time.Second
		time.Sleep(duration)
	}
}

