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

type transactionRepositories struct {
	elasticClient *elasticsearch8.TypedClient
	index         string
}

func NewTransactionRepositories(ctx context.Context, elasticClient *elasticsearch8.TypedClient, index string) ports.TransactionRepositories {
	res, err := elasticClient.Indices.Exists(index).Do(ctx)
	if err != nil {
		logger.Fatal("check indice error", logger.ErrField(err))
	}
	if !res {
		_, err := elasticClient.Indices.Create(index).Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"hex":      types.NewKeywordProperty(),
					"sender":   types.NewKeywordProperty(),
					"reciever": types.NewKeywordProperty(),
					"value":    types.NewUnsignedLongNumberProperty(),
					"gas":      types.NewUnsignedLongNumberProperty(),
					"gasPrice": types.NewUnsignedLongNumberProperty(),
					"nonce":    types.NewUnsignedLongNumberProperty(),
				},
			},
		}).Do(ctx)
		if err != nil {
			logger.Fatal("create indice with mapping error")
		}
	}
	return transactionRepositories{elasticClient: elasticClient, index: index}
}

func (t transactionRepositories) Save(ctx context.Context, txn domains.Transaction) (domains.Transaction, error) {
	res, err := t.elasticClient.Index(t.index).Id(txn.Hex).Request(txn).Do(ctx)
	if err != nil {
		return domains.Transaction{}, err
	}

	logger.Info("save transaction", logger.Field("response", res))
	return txn, nil
}
func (t transactionRepositories) GetByContainAddress(ctx context.Context, addr string, page int, perpage int) ([]domains.Transaction, error) {
	if page < 1 {
		page = 1
	}
	if perpage < 1 {
		perpage = 1
	}
	size := perpage
	from := (page - 1) * perpage

	res, err := t.elasticClient.Search().Index(t.index).Request(
		&search.Request{
			Query: &types.Query{
				Bool: &types.BoolQuery{
					Should: []types.Query{
						{
							Term: map[string]types.TermQuery{
								"sender": {Value: addr},
							},
						},
						{
							Term: map[string]types.TermQuery{
								"reciever": {Value: addr},
							},
						},
					},
				},
				// MatchAll: &types.MatchAllQuery{},
			},
			Size: &size,
			From: &from,
		},
	).Do(ctx)

	if err != nil {
		return []domains.Transaction{}, err
	}

	txns := []domains.Transaction{}
	for _, v := range res.Hits.Hits {
		b, err := v.Source_.MarshalJSON()
		if err != nil {
			return []domains.Transaction{}, err
		}
		tx := new(domains.Transaction)
		err = jsoniter.Unmarshal(b, tx)
		if err != nil {
			return []domains.Transaction{}, err
		}
		txns = append(txns, *tx)
	}

	return txns, nil
}
