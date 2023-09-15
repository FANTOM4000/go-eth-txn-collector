package main

import (
	httpserver "app/cmd/http_server"
	"app/config"
	"app/infrastructures"
	"app/internal/core/services"
	"app/internal/handler"
	"app/internal/repositories"
	"app/internal/repositories/adaptor"
	"app/pkg/logger"
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"
)

func main() {
	ctx := context.Background()
	config.Init()
	logger.Initlogger()
	client := resty.New()
	elasticsearchClient := infrastructures.InitElasticSearch(ctx)

	blockAdaptorApi := adaptor.NewBlockAdaptorApiRepositories(client, config.Get().BlockAdaptorApi.Url)
	addressRepositories := repositories.NewAddressRepositories(ctx, elasticsearchClient, config.Get().ElasticSearch.Index)
	addressService := services.NewAddressService(addressRepositories, blockAdaptorApi)
	addressHandler := handler.NewAddressHandler(addressService)
	app := httpserver.NewHttpServer(addressHandler)
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
