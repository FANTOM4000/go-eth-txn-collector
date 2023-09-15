package infrastructures

import (
	"app/config"
	"app/pkg/logger"
	"context"
	"strings"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

func InitElasticSearch(ctx context.Context) *elasticsearch8.TypedClient {
	es8, err := elasticsearch8.NewTypedClient(elasticsearch8.Config{
		Addresses: strings.Split(config.Get().ElasticSearch.Address, ","),
		Username:  config.Get().ElasticSearch.Username,
		Password:  config.Get().ElasticSearch.Password,
		
	})

	if err != nil {
		logger.Error("Cannot init elastic merchant client", logger.ErrField(err))
	}

	_, err = es8.Info().Do(ctx)
	if err != nil {
		logger.Panic("Cannot get elastic merchant info")
	}

	return es8
}
