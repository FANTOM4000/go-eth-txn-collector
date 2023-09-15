package adaptor

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"fmt"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type transactionAdaptorApiRepositories struct {
	restyClient *resty.Client
	Url         string
}

func NewTransactionAdaptorApiRepositories(restyClient *resty.Client, Url string) ports.TransactionAdaptorApiRepositories {
	return transactionAdaptorApiRepositories{restyClient: restyClient, Url: Url}
}

func (b transactionAdaptorApiRepositories) SaveTransaction(txn domains.Transaction) (domains.TransactionAdaptorApiResponse, error) {
	blockRes := new(domains.TransactionAdaptorApiResponse)
	resp, err := b.restyClient.R().
		SetBody(txn).
		Post(fmt.Sprint(b.Url,"/transaction"))
	jsoniter.Unmarshal(resp.Body(), &blockRes)

	if err != nil {
		return domains.TransactionAdaptorApiResponse{}, err
	}

	return *blockRes, nil
}
