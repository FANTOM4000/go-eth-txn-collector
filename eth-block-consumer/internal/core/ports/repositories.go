package ports

import "app/internal/core/domains"

type BlockAdaptorApiRepositories interface {
	ProduceTransaction(number uint64) (domains.BlockAdaptorApiResponse, error)
}
