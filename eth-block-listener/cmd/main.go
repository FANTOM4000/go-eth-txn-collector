package main

import (
	"app/config"
	"app/infrastructure"
	"app/internal/core/services"
	"app/internal/handler"
	"app/internal/repositories"
	"app/pkg/logger"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	config.Init()
	logger.Initlogger()
	ethClient := infrastructure.InitWebsockerEthereumClient()
	producer := infrastructure.InitKafkaProducer()
	defer producer.Close()
	blockQueRepo := repositories.NewBlockQueRepositories(producer,config.Get().Kafka.Topic)
	blockQueService := services.NewBlockQueService(blockQueRepo)
	blockQueHandler := handler.NewBlockQueHandler(ethClient, blockQueService)

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	go func() {
		<-sigchan
		fmt.Println("Received interrupt signal, shutting down listener...")
		producer.Close()
		os.Exit(0)
	}()

	blockQueHandler.ListenNewBlock()
}
