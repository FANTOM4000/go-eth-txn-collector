package main

import (
	httpserver "app/cmd/http_server"
	"app/config"
	"app/infrastructures"
	"app/internal/core/services"
	"app/internal/handler"
	"app/internal/repositories"
	"app/pkg/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx := context.Background()
	config.Init()
	logger.Initlogger()

	elasticsearchClient := infrastructures.InitElasticSearch(ctx)
	addressRepositoties := repositories.NewAddressRepositories(ctx, elasticsearchClient, config.Get().ElasticSearch.AddressIndex)
	transactionRepositories := repositories.NewTransactionRepositories(ctx, elasticsearchClient, config.Get().ElasticSearch.Index)
	transactionService := services.NewTransactionService(transactionRepositories, addressRepositoties)
	transactionHandler := handler.NewTransactionHandler(transactionService)
	app := httpserver.NewHttpServer(transactionHandler)
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
