package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Httpserver    httpserver
	ElasticSearch elasticSearch
}

type httpserver struct {
	Port string `envconfig:"HTTP_PORT" default:"80"`
}

type elasticSearch struct {
	Address  string `envconfig:"ELASTIC_ADDR" default:"localhost:9200"`
	Username string `envconfig:"ELASTIC_USERNAME" default:"username"`
	Password string `envconfig:"ELASTIC_PASSWORD" default:"password"`
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
