package httpserver

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewHttpServer(addressHandler ports.AddressHandler) *fiber.App {
	r := fiber.New()
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))

	addressGroup := r.Group("address")

	addressGroup.Post("/", addressHandler.AddAddressToWatch)
	addressGroup.Get("/", addressHandler.GetAllAddress)
	addressGroup.Delete("/{id}",addressHandler.Delete)
	return r
}
