package http

import (
	"evermos-project/internal/repository"
	"evermos-project/internal/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(productUsecase *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func NewProductHandler(productUsecase *usecase.ProdctUsecase) *ProductHandler {
	return &ProductHandler{productUsecase: productUsecase}
}

func (h *ProductHandler) GetProducts(c *fiber.Ctx) error {
	query := repository.ProductQuery{
		Search: c.Query("search"),
		Limit:  parseIntParam(c.Query("limit", "10")),
		Page:   parseIntParam(c.Query("page", "1")),
	}

	if categoryID := parseIntParam(c.Query("category_id")); categoryID > 0 {
		query.CategoryID = uint(categoryID)
	}
	if storeID := parseIntParam(c.Query("store_id")); storeID > 0 {
		query.StoreID = uint(storeID)
	}
	if minPrice := parseFloatparam(c.Query("min_price")); minPrice > 0 {
		query.MinPrice = minPrice
	}
	if maxPrice := parseFloatParam(c.Query("max_price")); maxPrice > 0 {
		query.MaxPrice = maxPrice
	}

	result, err := h.productUsecase.GetProducts(query)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(result)
}

func parseIntParam(param string, defaultVal ...string) int {
	if param == "" && len(defaultVal) > 0 {
		param = defaultVal[0]
	}
	val, _ := strconv.Atoi(param)
	return val
}

func parseFloatParam(param string) float64 {
	val, _ := strconv.ParseFloat(param, 64)
	return val
}
