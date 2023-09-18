package repositories

import (
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type addressRepositories struct {
	elasticClient *elasticsearch8.TypedClient
	index         string
}

func NewAddressRepositories(ctx context.Context, elasticClient *elasticsearch8.TypedClient, index string) ports.AddressRepositories {
	res, err := elasticClient.Indices.Exists(index).Do(ctx)
	if err != nil {
		logger.Fatal("check indice error", logger.ErrField(err))
	}
	if !res {
		_, err := elasticClient.Indices.Create(index).Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"hex": types.NewKeywordProperty(),
				},
			},
		}).Do(ctx)
		if err != nil {
			logger.Fatal("create indice with mapping error",logger.ErrField(err))
		}
	}

	return addressRepositories{elasticClient: elasticClient, index: index}
}

func (a addressRepositories) CheckAddressExist(ctx context.Context, addr string) (bool, error) {
	found, err := a.elasticClient.Exists(a.index, addr).Do(ctx)
	if err != nil {
		return false, err
	}
	return found, nil
}
