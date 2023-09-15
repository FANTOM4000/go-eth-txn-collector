package ports

import (
	"app/internal/core/domains"
	"context"
)

type TransactionRepositories interface {
	Save(ctx context.Context, txn domains.Transaction) (domains.Transaction, error)
	GetByContainAddress(ctx context.Context, addr string, page int, perpage int) ([]domains.Transaction, error)
}

type AddressRepositories interface {
	CheckAddressExist(ctx context.Context,addr string)(bool,error)
}