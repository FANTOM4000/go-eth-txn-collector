package ports

import "github.com/gofiber/fiber/v2"

type TransactionHandler interface {
	SaveTransactionData(ctx *fiber.Ctx) error
	GetAllTransactionOfAddress(ctx *fiber.Ctx) error
}
