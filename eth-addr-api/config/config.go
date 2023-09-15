package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Httpserver      httpserver
	ElasticSearch   elasticSearch
	BlockAdaptorApi blockAdaptorApi
}

type httpserver struct {
	Port string `envconfig:"HTTP_PORT" default:"80"`
}

type elasticSearch struct {
	Address  string `envconfig:"ELASTIC_ADDR" default:"localhost:9200"`
	Username string `envconfig:"ELASTIC_USERNAME" default:"username"`
	Password string `envconfig:"ELASTIC_PASSWORD" default:"password"`
	Index    string `envconfig:"ELASTIC_INDEX" default:"transaction"`
}

type blockAdaptorApi struct {
	Url string `envconfig:"BLOCK_ADAPTOR_API" default:"http://localhost:8080"`
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
