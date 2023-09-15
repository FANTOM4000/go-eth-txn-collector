package handler

import (
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type botHandler struct {
	blockService ports.BlockService
}

func NewBlockHandler(blockService ports.BlockService) ports.BlockHandler {
	return botHandler{blockService: blockService}
}

func (b botHandler) ProduceTrasactionFromBlockHash(ctx *fiber.Ctx) error {
	h := ctx.Params("hex")
	res, err := b.blockService.ProduceTrasactionFromBlockHash(ctx.Context(), h)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&res)
}
