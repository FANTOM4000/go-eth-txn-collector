package main

import (
	"app/config"
	"app/infrastructure"
	"app/pkg/logger"
	"fmt"
	"os"
	"os/signal"
)

func main() {

	config.Init()
	logger.Initlogger()

	// Create Kafka consumer
	consumer := infrastructure.InitKafkaConsumer()
	defer consumer.Close()

	

	// Trap signals to gracefully shut down the consumer
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	<-sigchan
	fmt.Println("Received interrupt signal, shutting down consumer...")
	consumer.Close()
	os.Exit(0)

}
