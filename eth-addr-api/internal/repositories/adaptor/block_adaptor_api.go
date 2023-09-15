package adaptor

import (
	"app/internal/core/domains"
	"app/internal/core/ports"
	"fmt"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type blockAdaptorApiRepositories struct {
	restyClient *resty.Client
	Url         string
}

func NewBlockAdaptorApiRepositories(restyClient *resty.Client, Url string) ports.BlockAdaptorApiRepositories {
	return blockAdaptorApiRepositories{restyClient: restyClient, Url: Url}
}

func (b blockAdaptorApiRepositories) ProduceTransaction(number uint64) (domains.BlockAdaptorApiResponse, error) {
	blockRes := new(domains.BlockAdaptorApiResponse)
	resp, err := b.restyClient.R().
		Post(fmt.Sprint(b.Url, "/block/number/", number, "/transactions"))
	jsoniter.Unmarshal(resp.Body(), &blockRes)

	if err != nil {
		return domains.BlockAdaptorApiResponse{}, err
	}

	return *blockRes, nil
}
