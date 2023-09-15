package ports

import "app/internal/core/domains"

type ConsumerService interface {
	Consume(txn domains.Transaction) error
}
