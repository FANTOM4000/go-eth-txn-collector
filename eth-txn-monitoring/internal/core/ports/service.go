package ports

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"context"
)

type AddressService interface {
	AddAddressToWatch(ctx context.Context, req dto.AddAddressToWatchRequest) (dto.BaseOKResponse, error)
	GetAll(ctx context.Context, page, perpage int) (dto.BaseOKResponseWithData[[]domains.Address], error)
	Delete(ctx context.Context, id string) (dto.BaseOKResponse, error)
}

type TransactionService interface {
	Save(ctx context.Context, tnx domains.Transaction) (dto.BaseOKResponseWithData[domains.Transaction], error)
	GetIncomingAndOutgoingOfAddress(ctx context.Context, addr string, page int, perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error)
}
