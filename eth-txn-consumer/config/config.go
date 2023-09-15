package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Kafka                 kafka
	TransactionAdaptorApi transactionAdaptorApi
}
type kafka struct {
	Server   string `envconfig:"KAFKA_SERVER" default:"localhost:9092"`
	Group    string `envconfig:"KAFKA_GROUP" default:"demo_group"`
	Topic    string `envconfig:"KAFKA_TOPIC" default:"demo.topic"`
	Username string `envconfig:"KAFKA_USERNAME" default:"username"`
	Password string `envconfig:"KAFKA_PASSWORD" default:"password"`
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
