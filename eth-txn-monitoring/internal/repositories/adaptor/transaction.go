package adaptor

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"context"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type transactionAdaptorApi struct {
	restyClient *resty.Client
	Url         string
}

func NewTransactionAdaptorApi(restyClient *resty.Client, Url string) ports.TransactionAdaptorApi {
	return transactionAdaptorApi{restyClient: restyClient, Url: Url}
}

func (t transactionAdaptorApi) GetIncomingAndOutgoingOfAddress(ctx context.Context, addr string, page int, perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error) {
	res, err := t.restyClient.R().SetQueryParam("addr", addr).SetQueryParam("page", strconv.Itoa(page)).SetQueryParam("perpage", strconv.Itoa(perpage)).Get(fmt.Sprint(t.Url,"/transaction"))
	if err != nil {
		return dto.BaseOKResponseWithData[[]domains.Transaction]{}, err
	}
	response := dto.BaseOKResponseWithData[[]domains.Transaction]{}
	jsoniter.Unmarshal(res.Body(), &response)
	return response, nil
}
