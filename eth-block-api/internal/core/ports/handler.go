package ports

import "github.com/gofiber/fiber/v2"

type BlockHandler interface {
	ProduceTrasactionFromBlockHash(ctx *fiber.Ctx) error
}