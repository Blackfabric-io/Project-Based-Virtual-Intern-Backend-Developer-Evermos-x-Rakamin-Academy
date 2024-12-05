package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	CategoryID  uint    `json:"category_id"`
	StoreID     uint    `json:"store_id"`
}

type ProductRepository interface {
	GetProducts(query *ProductQuery) ([]Product, int64, error)
	GetProductByID(id uint) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
}

type ProductUsecase interface {
	GetProducts(query *ProductQuery) (*PaginatedResponse, error)
	GetProductByID(id uint) (*Product, error)
	CreateProduct(product *Product) error
	UpdateProduct(product *Product) error
	DeleteProduct(id uint) error
}

type ProductQuery struct {
	Search     string
	CategoryID uint
	StoreID    uint
	MinPrice   float64
	MaxPrice   float64
	Page       int
	Limit      int
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	CurrentPage  int   `json:"current_page"`
	TotalPages   int   `json:"total_pages"`
	TotalRecords int64 `json:"total_records"`
	Limit        int   `json:"limit"`
}
