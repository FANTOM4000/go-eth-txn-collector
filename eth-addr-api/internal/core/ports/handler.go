package ports

import "github.com/gofiber/fiber/v2"

type AddressHandler interface {
	AddAddressToWatch(ctx *fiber.Ctx) error
	GetAllAddress(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
