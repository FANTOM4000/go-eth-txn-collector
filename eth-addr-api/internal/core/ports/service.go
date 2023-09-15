package ports

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"context"
)

type AddressService interface {
	AddAddressToWatch(ctx context.Context, addr domains.Address, fromBlock int, toBlock int) (dto.BaseOKResponse, error)
	GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error)
	Delete(ctx context.Context, id string) (dto.BaseOKResponse, error)
}
