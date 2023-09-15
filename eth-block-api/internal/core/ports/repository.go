package ports

import (
	"app/internal/core/domains"
	"context"
)

type EthRepositories interface {
	GetTransactionByBlockHash(ctx context.Context,number uint64) ([]domains.Transaction, error)
}

type TransactionQueRepositories interface {
	Produce(domains.Transaction) error
}
