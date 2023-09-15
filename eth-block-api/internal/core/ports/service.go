package ports

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
	"context"
)

type BlockService interface {
	ProduceTrasactionFromBlockHash(ctx context.Context, hex string) (dto.BaseOKResponseWithData[[]domains.Transaction], error)
}
