package ports

import "github.com/gofiber/fiber/v2"

type AddressHandler interface {
	AddAddressToWatch(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type TransactionHandler interface {
	Save(ctx *fiber.Ctx) error
	GetIncomingAndOutgoingOfAddress(ctx *fiber.Ctx) error
}
