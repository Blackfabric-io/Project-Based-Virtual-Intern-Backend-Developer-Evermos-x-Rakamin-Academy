package usecase

import {
	"evermoss-project/internal/domain"
	"evermoss-project/internal/repository"
	"math"
}

type ProductUseCase struct {
	productRepo *repository.ProductRepository
}

func NewProductUsecase(productRepo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase(productRepo: productRepo)
}

type PaginatedResponse struct {
	Data	interface{} `json:"data"`
    Pagination Pagination  `json:"pagination"
}

type Pagination struct {
	CurentPage int `json:"current_page"`
	TotalPages int `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
	Limit int `json:"limit"`
}

func (u *ProductUseCase) GetProducts(query repository.ProductQuery) (*PaginatedResponse, error) {
	products, total, err := u.productRepo.GetProducts(query)
	if err != nil {
		return nil, err
	}

	totalPages := int(math.Ceil(float64(total) / float64(query.Limit)))

	return &PaginatedResponse{
		Data: products,
		Pagination: Pagination{
			CurentPage: query.Page,
			TotalPages: totalPages,
			TotalRecords: total,
			Limit: query.Limit,
		},
	}, nil
}


