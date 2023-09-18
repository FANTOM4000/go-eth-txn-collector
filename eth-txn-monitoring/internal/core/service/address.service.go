package service

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"context"
)

type addressService struct {
	addressAdaptor ports.AddressAdaptorApi
}

func NewAddressService(addressAdaptor ports.AddressAdaptorApi) ports.AddressService {
	return addressService{addressAdaptor: addressAdaptor}
}

func (a addressService) AddAddressToWatch(ctx context.Context, req dto.AddAddressToWatchRequest) (dto.BaseOKResponse, error){
	return a.addressAdaptor.AddAddressToWatch(ctx,req)
}
func (a addressService) GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error){
	return a.addressAdaptor.GetAll(ctx,page,perpage)
}
func (a addressService) Delete(ctx context.Context, id string) (dto.BaseOKResponse, error) {
	return a.addressAdaptor.Delete(ctx,id)
}
