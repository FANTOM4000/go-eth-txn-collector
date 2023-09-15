package ports

import "app/internal/core/domains"

type TransactionAdaptorApiRepositories interface {
	SaveTransaction(txn domains.Transaction) (domains.TransactionAdaptorApiResponse, error)
}
