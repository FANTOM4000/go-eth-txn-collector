package main

import (
	httpserver "app/cmd/http_server"
	"app/config"
	"app/internal/core/service"
	"app/internal/handler"
	"app/internal/repositories/adaptor"
	"app/pkg/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-resty/resty/v2"
)

func main() {
	config.Init()
	logger.Initlogger()
	client := resty.New()

	addressAdaptor := adaptor.NewAddressAdaptorApi(client, config.Get().AddressAdaptorApi.Url)
	transactonAdaptor := adaptor.NewTransactionAdaptorApi(client, config.Get().TransactionAdaptorApi.Url)
	addressService := service.NewAddressService(addressAdaptor)
	transactionService := service.NewTransactionService(transactonAdaptor)
	addressHandler := handler.NewAddressHandler(addressService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	app := httpserver.NewHttpServer(addressHandler, transactionHandler)
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
