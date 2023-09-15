package ports

import (
	"app/internal/core/domains"
	"context"
)

type AddressRepositories interface {
	Create(ctx context.Context, addr domains.Address) (string, error)
	GetById(ctx context.Context, id string) (domains.Address, error)
	GetAll(ctx context.Context, page, perpage int) ([]domains.Address, error)
	Delete(ctx context.Context, id string) error
}

type BlockAdaptorApiRepositories interface {
	ProduceTransaction(number uint64) (domains.BlockAdaptorApiResponse, error)
}
