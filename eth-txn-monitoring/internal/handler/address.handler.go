package handler

import (
	"app/internal/core/dto"
	"app/internal/core/ports"

	"github.com/gofiber/fiber/v2"
)

type addressHandler struct {
	addressService ports.AddressService
}

func NewAddressHandler(addressService ports.AddressService) ports.AddressHandler {
	return addressHandler{addressService: addressService}
}

func (a addressHandler) AddAddressToWatch(ctx *fiber.Ctx) error {
	req := dto.AddAddressToWatchRequest{}
	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.Status(400).JSON(nil)
	}
	res, err := a.addressService.AddAddressToWatch(ctx.Context(), req)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
func (a addressHandler) GetAll(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	perpage := ctx.QueryInt("perpage", 1)
	res, err := a.addressService.GetAll(ctx.Context(), page, perpage)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
func (a addressHandler) Delete(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	res, err := a.addressService.Delete(ctx.Context(), id)
	if err != nil {
		return ctx.Status(500).JSON(res)
	}
	return ctx.JSON(res)
}
