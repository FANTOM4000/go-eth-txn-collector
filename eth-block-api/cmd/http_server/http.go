package httpserver

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHttpServer(blockHandler ports.BlockHandler) *fiber.App {
	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	r.Post("/block/{hex}/transactions",blockHandler.ProduceTrasactionFromBlockHash)

	return r
}
