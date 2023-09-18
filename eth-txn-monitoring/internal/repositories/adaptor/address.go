package adaptor

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"app/pkg/logger"
	"context"
	"fmt"
	"strconv"

	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
)

type addressAdaptorApi struct {
	restyClient *resty.Client
	Url         string
}

func NewAddressAdaptorApi(restyClient *resty.Client, Url string) ports.AddressAdaptorApi {
	return addressAdaptorApi{restyClient: restyClient, Url: Url}
}

func (a addressAdaptorApi) AddAddressToWatch(ctx context.Context, req dto.AddAddressToWatchRequest) (dto.BaseOKResponse, error) {
	res, err := a.restyClient.R().SetBody(req).Post(fmt.Sprint(a.Url,"/address"))
	if err != nil {
		return dto.BaseOKResponse{}, err
	}
	response := dto.BaseOKResponse{}
	jsoniter.Unmarshal(res.Body(), &response)
	return response, nil
}
func (a addressAdaptorApi) GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error) {
	res, err := a.restyClient.R().SetQueryParam("page", strconv.Itoa(page)).SetQueryParam("perpage", strconv.Itoa(perpage)).Get(fmt.Sprint(a.Url,"/address"))
	if err != nil {
		logger.Error("fail call GetAll ", logger.ErrField(err))
		return dto.BaseOKResponseWithData[[]domains.Address]{}, err
	}
	response := dto.BaseOKResponseWithData[[]domains.Address]{}
	jsoniter.Unmarshal(res.Body(), &response)
	logger.Info("response from GetAll",logger.Field("response",response))
	return response, nil
}
func (a addressAdaptorApi) Delete(ctx context.Context, id string) (dto.BaseOKResponse, error) {
	res, err := a.restyClient.R().Delete(fmt.Sprint(a.Url, "/address/", id))
	if err != nil {
		return dto.BaseOKResponse{}, err
	}
	response := dto.BaseOKResponse{}
	jsoniter.Unmarshal(res.Body(), &response)
	return response, nil
}
