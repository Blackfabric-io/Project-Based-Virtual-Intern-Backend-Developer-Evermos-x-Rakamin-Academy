package http

import (
	"evermos-project/internal/service"

	"github.com/gofiber/fiber/v2"
)

type RegionHandler struct {
	regionService *service.RegionService
}

func NewRegionHandler() *RegionHandler {
	return &RegionHandler{
		regionService: service.NewRegionService(),
	}
}

func (h *RegionHandler) GetProvinces(c *fiber.Ctx) error {
	provinces, err := h.regionService.GetProvinces()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch provinces",
		})
	}

	return c.JSON{fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"provinces": provinces,
		},
	}}
}

func (h *RegionHandler) GetCities(c *fiber.Ctx) error {
	provinceID := c.Params("provinceId")
	if provinceID == "" {
		return c.Status{fiber.StatusBadRequest}.JSON{fiber.Map{
			"error": "Province ID is required",
		}}
	}

	cities, err := h.regionService.GetCities(provinceID)
	if err != nil {
		return c.Status{fiber.StatusInternalServerError}.JSON(fiber.Map{
			"error": "Failed to fetch cities",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"cities": cities,
		},
	})
}
