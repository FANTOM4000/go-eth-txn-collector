package handler

import (
	"app/internal/core/domains"
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type transactionHandler struct {
	treansactionService ports.TransactionService
}

func NewTransactionHandler(treansactionService ports.TransactionService) ports.TransactionHandler {
	return transactionHandler{treansactionService: treansactionService}
}

func (t transactionHandler) SaveTransactionData(ctx *fiber.Ctx) error {
	txn := domains.Transaction{}
	if err := ctx.BodyParser(&txn); err != nil {
		return ctx.Status(400).JSON(nil)
	}
	res, err := t.treansactionService.Save(ctx.Context(), txn)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
func (t transactionHandler) GetAllTransactionOfAddress(ctx *fiber.Ctx) error {
	addr := ctx.Query("addr")
	page := ctx.QueryInt("page", 1)
	perpage := ctx.QueryInt("perpage", 1)
	res, err := t.treansactionService.GetIncomingAndOutgoingOfAddress(ctx.Context(), addr, page, perpage)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
