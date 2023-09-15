package adaptor

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"context"

	"github.com/go-resty/resty/v2"
)

type addressAdaptorApi struct {
	restyClient *resty.Client
	Url         string
}

func NewAddressAdaptorApi(restyClient *resty.Client, Url string) ports.AddressAdaptorApi {
	return addressAdaptorApi{restyClient: restyClient, Url: Url}
}

func (a addressAdaptorApi) AddAddressToWatch(ctx context.Context, req dto.AddAddressToWatchRequest) (dto.BaseOKResponse, error) {

}
func (a addressAdaptorApi) GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error)
func (a addressAdaptorApi) Delete(ctx context.Context, id string) (dto.BaseOKResponse, error)
