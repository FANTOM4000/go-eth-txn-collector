package ports

import (
	"app/internal/core/domains"
	"app/internal/core/dto"
)

type TransactionService interface {
	Save(tnx domains.Transaction) (dto.BaseOKResponseWithData[domains.Transaction], error)
	GetIncomingAndOutgoingOfAddress(addr string,page int,perpage int) (dto.BaseOKResponseWithData[[]domains.Transaction], error)
}
