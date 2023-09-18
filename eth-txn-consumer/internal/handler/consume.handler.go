package handler

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/logger"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	jsoniter "github.com/json-iterator/go"
)

type consumeHandler struct {
	consumer       *kafka.Consumer
	consumeService ports.ConsumerService
	Topic          string
}

func NewConsumerHandler(consumer *kafka.Consumer, Topic string, consumeService ports.ConsumerService) ports.ConsumeHandler {
	return consumeHandler{consumer: consumer, consumeService: consumeService, Topic: Topic}
}

func (c consumeHandler) Consume() {

	topic := c.Topic
	err := c.consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		panic(fmt.Errorf("failed to subscribe topics : %s", err.Error()))
	}
	// Set up a channel for handling Ctrl-C, etc
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Process messages
	log.Printf("consumer is running...\n")
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("caught signal %v: terminating\n", sig)
			run = false
		default:
			ev, err := c.consumer.ReadMessage(100 * time.Millisecond)
			if err != nil {
				// Errors are informational and automatically handled by the consumer
				continue
			}

			log.Printf("consumed topic %s: partition: %d message: %s\n", *ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.Value)
			switch *ev.TopicPartition.Topic {
			case c.Topic:
				txn := domains.Transaction{}
				err = jsoniter.UnmarshalFromString(string(ev.Value), &txn)
				if err != nil {
					logger.Error("error unmarshal json", logger.ErrField(err))
					return
				}
				err = c.consumeService.Consume(txn)
				if err != nil {
					logger.Panic("error consume", logger.ErrField(err))
					return
				}
			}

			c.consumer.Commit()
		}
	}
}
