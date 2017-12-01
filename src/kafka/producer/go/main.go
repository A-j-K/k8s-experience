package main

import (
	"fmt"
	"os"
	"time"
	"github.com/Shopify/sarama"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	broker := os.Getenv("KAFKA_BROKER")
	fmt.Printf("Broker: %s\n", broker)
	brokers := []string{ broker }
	topic := os.Getenv("KAFKA_TOPIC")
	fmt.Printf("Topic: %s\n", topic)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
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
			Topic: topic,
			Value: sarama.StringEncoder(name),
		}
		partition, offset, err := producer.SendMessage(msg)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
		duration := time.Second
		time.Sleep(duration)
	}
}

