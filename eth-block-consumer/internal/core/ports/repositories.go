package ports

import "app/internal/core/domains"

type BlockAdaptorApiRepositories interface {
	ProduceTransaction(hash string) (domains.BlockAdaptorApiResponse, error)
}
