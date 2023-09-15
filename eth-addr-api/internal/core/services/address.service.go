package services

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"app/internal/core/ports"
	"app/pkg/standard"
	"context"
)

type addressService struct {
	addressRepositories ports.AddressRepositories
	blockAdaptorApi     ports.BlockAdaptorApiRepositories
}

func NewAddressService(addressRepositories ports.AddressRepositories, blockAdaptorApi ports.BlockAdaptorApiRepositories) ports.AddressService {
	return addressService{addressRepositories: addressRepositories, blockAdaptorApi: blockAdaptorApi}
}

func (a addressService) AddAddressToWatch(ctx context.Context, addr domains.Address, fromBlock int, toBlock int) (dto.BaseOKResponse, error) {
	_, err := a.addressRepositories.Create(ctx, addr)
	if err != nil {
		return dto.BaseOKResponse{
			Code:    standard.CreateError,
			Message: "add address error",
		}, err
	}
	if fromBlock > 0 || toBlock > 0 {
		min := fromBlock
		max := toBlock
		if min > max {
			tmp := min
			min = max
			max = tmp
		}
		for i := min; i < max; i++ {
			_, err := a.blockAdaptorApi.ProduceTransaction(uint64(i))
			if err != nil {
				return dto.BaseOKResponse{
					Code:    standard.GenericError,
					Message: "produce block to collect txn error",
				}, err
			}
		}
	}
	return dto.BaseOKResponse{
		Code:    standard.SuccessCode,
		Message: "success",
	}, nil
}
func (a addressService) GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error) {
	addrs, err := a.addressRepositories.GetAll(ctx, page, perpage)
	if err != nil {
		return dto.BaseOKResponseWithData[[]domains.Address]{
			BaseOKResponse: dto.BaseOKResponse{
				Code:    standard.GetDataError,
				Message: "get address error",
			},
			Data: []domains.Address{},
		}, err
	}

	return dto.BaseOKResponseWithData[[]domains.Address]{
		BaseOKResponse: dto.BaseOKResponse{
			Code:    standard.SuccessCode,
			Message: "success",
		},
		Data: addrs,
	}, nil
}
func (a addressService) Delete(ctx context.Context, id string) (dto.BaseOKResponse, error) {
	err := a.addressRepositories.Delete(ctx, id)
	if err != nil {
		return dto.BaseOKResponse{
			Code:    standard.DeleteError,
			Message: "delete error",
		}, err
	}
	return dto.BaseOKResponse{
		Code:    standard.SuccessCode,
		Message: "success",
	}, nil
}
