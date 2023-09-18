package handler

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	transactionService ports.TransactionService
}

func NewTransactionHandler(transactionService ports.TransactionService) ports.TransactionHandler {
	return transactionHandler{transactionService: transactionService}
}

func (t transactionHandler) GetIncomingAndOutgoingOfAddress(ctx *fiber.Ctx) error {
	addr := ctx.Query("adr")
	page := ctx.QueryInt("page", 1)
	perpage := ctx.QueryInt("perpage", 1)
	res, err := t.transactionService.GetIncomingAndOutgoingOfAddress(ctx.Context(), addr, page, perpage)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
