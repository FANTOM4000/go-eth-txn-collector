package repositories

import (
	"app/internal/core/ports"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type blockQueRepositories struct {
	kafkaClient *kafka.Producer
	Topic       string
}

func NewBlockQueRepositories(kafkaClient *kafka.Producer, Topic string) ports.BlockQueRepositories {
	return blockQueRepositories{kafkaClient: kafkaClient, Topic: Topic}
}

func (b blockQueRepositories) Produce(hex string) error {
	deliverChan := make(chan kafka.Event)

	b.kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &b.Topic, Partition: -1},
		Key:            []byte(hex),
		Value:          []byte(hex),
	}, deliverChan)
	<-deliverChan
	return nil
}