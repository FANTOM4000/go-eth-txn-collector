package main

import (
	httpserver "app/cmd/http_server"
	"app/config"
	"app/infrastructure"
	"app/internal/core/services"
	"app/internal/handler"
	"app/internal/repositories"
	"app/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config.Init()
	logger.Initlogger()

	producer := infrastructure.InitKafkaProducer()
	ethClient := infrastructure.InitEthereumClient()
	ethRepositories := repositories.NewEthRepositories(ethClient)
	transactionRepo := repositories.NewTransactionQueRepositories(producer, config.Get().Kafka.Topic)
	blockService := services.NewBlockService(ethRepositories, transactionRepo)
	blockHandler := handler.NewBlockHandler(blockService)
	app := httpserver.NewHttpServer(blockHandler)
	go func() {
		if err := app.Listen(fmt.Sprint(":", config.Get().Httpserver.Port)); err != nil {
			logger.Panic("listen server error")
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	<-c
	fmt.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	fmt.Println("Running cleanup tasks...")
	fmt.Println("Fiber was successful shutdown.")
}
