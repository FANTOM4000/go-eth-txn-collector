package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Httpserver      httpserver
	AddressAdaptorApi addressAdaptorApi
}

type httpserver struct {
	Port string `envconfig:"HTTP_PORT" default:"80"`
}

type addressAdaptorApi struct {
	Url string `envconfig:"ADDRESS_ADAPTOR_API" default:"http://localhost:8080"`
}

type transactionAdaptorApi struct {
	Url string `envconfig:"TRANSACTION_ADAPTOR_API" default:"http://localhost:8080"`
}

var cfg config

func Init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		err = godotenv.Load("../.env")
	}

	if err != nil {
		log.Printf("load env error : %s", err.Error())
	}
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatalf("read env error : %s", err.Error())
	}
}

func Get() config {
	return cfg
}
