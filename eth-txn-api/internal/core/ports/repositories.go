package ports

import "app/internal/core/domains"

type TransactionRepositories interface {
	Save(txn domains.Transaction) (domains.Transaction, error)
	GetByContainAddress(addr string, page int, perpage int) ([]domains.Transaction, error)
}
