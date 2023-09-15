package ports

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"context"
)

type TransactionService interface {
	Save(ctx context.Context, tnx domains.Transaction) (dto.BaseOKResponseWithData[domains.Transaction], error)
	GetIncomingAndOutgoingOfAddress(ctx context.Context, addr string, page int, perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error)
}
