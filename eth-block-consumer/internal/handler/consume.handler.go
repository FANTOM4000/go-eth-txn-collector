package handler

import (
	"app/internal/core/ports"
	"app/pkg/logger"
	"fmt"
	"strconv"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type consumeHandler struct {
	consumer       *kafka.Consumer
	consumeService ports.ConsumerService
	Topic          string
}

func NewConsumerHandler(consumer *kafka.Consumer, Topic string, consumeService ports.ConsumerService) ports.ConsumeHandler {
	return consumeHandler{consumer: consumer, consumeService: consumeService}
}

func (c consumeHandler) Consume() {
	topics := []string{c.Topic}
	err := c.consumer.SubscribeTopics(topics, nil)
	if err != nil {
		logger.Error("Error subscribing to topics", logger.ErrField(err))
		return
	}
	// Start consuming messages
	logger.Info("Kafka consumer started. Waiting for messages...")
	for {
		select {
		case ev := <-c.consumer.Events():
			switch e := ev.(type) {
			case *kafka.Message:
				logger.Info("Received message", logger.StringField("topic", *e.TopicPartition.Topic), logger.StringField("partition", e.TopicPartition.Partition), logger.StringField("offset", e.TopicPartition.Offset), logger.StringField("message", string(e.Value)))
				number,err := strconv.ParseUint(string(e.Value),10,64)
				if err != nil {
					logger.Error("error parse uint64", logger.ErrField(err))
					return
				}
				err = c.consumeService.Consume(number)
				if err != nil {
					logger.Error("error consume", logger.ErrField(err))
					return
				}
				c.consumer.Commit()
			case kafka.Error:
				fmt.Printf("Kafka error: %v\n", e)
			default:
				fmt.Printf("Ignored event: %v\n", e)
			}
		}
	}
}
