package handler

import (
	"app/internal/core/ports"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type botHandler struct {
	blockService ports.BlockService
}

func NewBlockHandler(blockService ports.BlockService) ports.BlockHandler {
	return botHandler{blockService: blockService}
}

func (b botHandler) ProduceTrasactionFromBlockHash(ctx *fiber.Ctx) error {
	nstr := ctx.Params("number")
	number,err := strconv.ParseUint(nstr,10,64)
	if err!= nil {
		ctx.Status(400).JSON(nil)
	}
	res, err := b.blockService.ProduceTrasactionFromBlockHash(ctx.Context(), number)
	if err != nil {
		return ctx.Status(500).JSON(nil)
	}
	return ctx.JSON(&res)
}
