package httpserver

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"
)

func NewHttpServer(addressHandler ports.AddressHandler, transactionHandler ports.TransactionHandler) *fiber.App {
	engine := django.New("./views", ".html")
	r := fiber.New(fiber.Config{
		Views: engine,
	})
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${time}] ${ip}  ${status} - ${latency} ${method} ${path}\n",
	}))
	r.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	r.Post("/add-ddress-to-watch", addressHandler.AddAddressToWatch)
	r.Get("/get-address-all", addressHandler.GetAll)
	r.Delete("/delete-address/{id}", addressHandler.Delete)

	r.Get("/all-transaction", transactionHandler.GetIncomingAndOutgoingOfAddress)

	return r
}
