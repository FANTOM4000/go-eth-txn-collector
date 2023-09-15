package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	jsoniter "github.com/json-iterator/go"
)

type transactionQueRepositories struct {
	kafkaClient *kafka.Producer
	Topic       string
}

func NewTransactionQueRepositories(kafkaClient *kafka.Producer, Topic string) ports.TransactionQueRepositories {
	return transactionQueRepositories{kafkaClient: kafkaClient, Topic: Topic}
}

func (t transactionQueRepositories) Produce(txn domains.Transaction) error {
	deliverChan := make(chan kafka.Event)
	b, err := jsoniter.Marshal(&txn)
	if err != nil {
		return err
	}
	t.kafkaClient.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &t.Topic, Partition: -1},
		Key:            []byte(txn.Hex),
		Value:          b,
	}, deliverChan)
	<-deliverChan
	return nil
}
