package repositories

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"app/pkg/logger"
	"bytes"
	"fmt"
	"strings"

	elasticsearch7 "github.com/elastic/go-elasticsearch/v7"
	jsoniter "github.com/json-iterator/go"
)

type transactionRepositories struct {
	elasticClient *elasticsearch7.Client
	index         string
}

func NewTransactionRepositories(elasticClient *elasticsearch7.Client, index string) ports.TransactionRepositories {
	return transactionRepositories{elasticClient: elasticClient, index: index}
}

func (t transactionRepositories) Save(txn domains.Transaction) (domains.Transaction, error) {
	b, err := jsoniter.Marshal(txn)
	if err != nil {
		return domains.Transaction{}, err
	}

	res, err := t.elasticClient.Index(t.index, bytes.NewReader(b), t.elasticClient.Index.WithDocumentID(txn.Hex))
	if err != nil {
		return domains.Transaction{}, err
	}
	logger.Info("save transaction", logger.StringField("response", res.String()))
	return txn, nil
}
func (t transactionRepositories) GetByContainAddress(addr string, page int, perpage int) ([]domains.Transaction, error) {
	query := fmt.Sprintf(`{"query":{"bool":{"should":[{"term":{"sender":"%s"}},{"term":{"reciever":"%s"}}]}}}`, addr,addr)
	t.elasticClient.Search(
		t.elasticClient.Search.WithIndex(t.index),
		t.elasticClient.Search.WithBody(strings.NewReader(query)),
	)
}
