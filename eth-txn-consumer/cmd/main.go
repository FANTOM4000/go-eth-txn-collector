package main

import (
	"app/config"
	"app/infrastructure"
	"app/internal/core/service"
	"app/internal/handler"
	"app/internal/repositories/adaptor"
	"app/pkg/logger"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-resty/resty/v2"
)

func main() {
	config.Init()
	logger.Initlogger()
	// Create Kafka consumer
	consumer := infrastructure.InitKafkaConsumer()
	defer consumer.Close()

	client := resty.New()

	blockAdaptorRepo := adaptor.NewTransactionAdaptorApiRepositories(client, config.Get().TransactionAdaptorApi.Url)
	consumeService := service.NewConsumeService(blockAdaptorRepo)
	consumeHandler := handler.NewConsumerHandler(consumer, config.Get().Kafka.Topic, consumeService)

	// Trap signals to gracefully shut down the consumer
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		<-sigchan
		fmt.Println("Received interrupt signal, shutting down consumer...")
		consumer.Close()
		os.Exit(0)
	}()

	consumeHandler.Consume()

}
