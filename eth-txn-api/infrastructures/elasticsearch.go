package infrastructures

import (
	"app/config"
	"app/pkg/logger"
	"fmt"
	"net/http"
	"strings"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

func InitElasticSearch() *elasticsearch8.Client {
	es8, err := elasticsearch8.NewClient(elasticsearch8.Config{
		Addresses: strings.Split(config.Get().ElasticSearch.Address, ","),
		Username:  config.Get().ElasticSearch.Username,
		Password:  config.Get().ElasticSearch.Password,
	})


	if err != nil {
		logger.Error("Cannot init elastic merchant client", logger.ErrField(err))
	}

	info, err := es8.Info()
	if err != nil {
		logger.Panic("Cannot get elastic merchant info")
	}

	if info.StatusCode != http.StatusOK {
		logger.Panic(fmt.Sprintf("Cannot connect to elasticsearch merchant, http_status=%d", info.StatusCode), logger.StringField("es_hosts", config.Get().ElasticSearch.Address), logger.StringField("es_username", config.Get().ElasticSearch.Username), logger.StringField("es_password", config.Get().ElasticSearch.Password))
	}

	logger.Info(fmt.Sprintf("Connect to elasticsearch merchant success, http_status=%d", info.StatusCode))
	return es8
}
