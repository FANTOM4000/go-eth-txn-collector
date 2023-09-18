package httpserver

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHttpServer(transactionHandler ports.TransactionHandler) *fiber.App {
	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))
	r.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("running")
	})
	transactionGroup := r.Group("transaction")

	transactionGroup.Post("/", transactionHandler.SaveTransactionData)
	transactionGroup.Get("/", transactionHandler.GetAllTransactionOfAddress)

	return r
}
