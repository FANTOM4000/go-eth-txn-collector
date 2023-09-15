package infrastructure

import (
	"app/config"
	"app/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func InitKafkaConsumer() *kafka.Consumer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers": config.Get().Kafka.Server, // Replace with your Kafka broker(s)
		"group.id":          config.Get().Kafka.Group,  // Replace with your consumer group name
		"auto.offset.reset": "earliest",                // Set the offset to the beginning (you can change to "latest" if you want the latest messages)
		"sasl.mechanism":    "PLAIN",
		"security.protocol": "SASL_PLAINTEXT",
		"sasl.username":     config.Get().Kafka.Username, // Replace with your username
		"sasl.password":     config.Get().Kafka.Password, // Replace with your password
	}
	consumer, err := kafka.NewConsumer(cfg)
	if err != nil {
		logger.Fatal("init kafka error", logger.ErrField(err))
	}
	return consumer
}

func InitKafkaProducer() *kafka.Producer {
	cfg := &kafka.ConfigMap{
		"bootstrap.servers":  config.Get().Kafka.Server, // Replace with your Kafka broker(s)
		"group.id":           config.Get().Kafka.Group,  // Replace with your consumer group name
		"auto.offset.reset":  "earliest",                // Set the offset to the beginning (you can change to "latest" if you want the latest messages)
		"enable.auto.commit": false,
		"sasl.mechanism":     "PLAIN",
		"security.protocol":  "SASL_PLAINTEXT",
		"sasl.username":      config.Get().Kafka.Username, // Replace with your username
		"sasl.password":      config.Get().Kafka.Password, // Replace with your password
	}

	producer, err := kafka.NewProducer(cfg)

	if err != nil {
		logger.Fatal("error new producer", logger.ErrField(err))
	}
	return producer
}
