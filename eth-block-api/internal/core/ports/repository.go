package ports

import (
	"app/internal/core/domains"
	"context"
)

type EthRepositories interface {
	GetTransactionByBlockHash(ctx context.Context, hex string) ([]domains.Transaction, error)
}

type TransactionQueRepositories interface {
	Produce(domains.Transaction) error
}
