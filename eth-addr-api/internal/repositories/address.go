package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"

	elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v8/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	jsoniter "github.com/json-iterator/go"
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
			logger.Fatal("create indice with mapping error")
		}
	}

	return addressRepositories{elasticClient: elasticClient, index: index}
}

func (a addressRepositories) Create(ctx context.Context, addr domains.Address) (string, error) {
	res, err := a.elasticClient.Index(a.index).Id(addr.Hex).Document(&addr).Do(ctx)
	if err != nil {
		return "", err
	}
	return res.Id_, nil
}
func (a addressRepositories) GetById(ctx context.Context, id string) (domains.Address, error) {
	res, err := a.elasticClient.Get(a.index, id).Do(ctx)
	if err != nil {
		return domains.Address{}, err
	}
	b, err := res.Source_.MarshalJSON()
	if err != nil {
		return domains.Address{}, err
	}
	addr := domains.Address{}
	jsoniter.Unmarshal(b, &addr)
	return addr, nil
}
func (a addressRepositories) GetAll(ctx context.Context, page, perpage int) ([]domains.Address, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 1
	}
	size := perpage
	from := (page - 1) * perpage

	res, err := a.elasticClient.Search().
		Index(a.index).
		Request(&search.Request{
			Query: &types.Query{
				MatchAll: &types.MatchAllQuery{},
			},
			Size: &size,
			From: &from,
		}).
		Do(ctx)
	if err != nil {
		return []domains.Address{}, err
	}

	addrs := []domains.Address{}
	for _, addrHit := range res.Hits.Hits {
		b, err := addrHit.Source_.MarshalJSON()
		if err != nil {
			return []domains.Address{}, err
		}
		addr := domains.Address{}
		err = jsoniter.Unmarshal(b, &addr)
		if err != nil {
			return []domains.Address{}, err
		}
		addrs = append(addrs, addr)
	}

	return addrs, nil
}
func (a addressRepositories) Delete(ctx context.Context, id string) error {
	_,err := a.elasticClient.Delete(a.index,id).Do(ctx)
	if err!=nil {
		return err
	}
	return nil
}
