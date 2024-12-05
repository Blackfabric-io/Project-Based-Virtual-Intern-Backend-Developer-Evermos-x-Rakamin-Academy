package repository

inport (
	"evermos-project/internal/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db &gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

type ProductQuery struct {
	Search string
	Limit int
	Page int
	CategoryID int
	StoreID int
	MinPrice float64
	MaxPrice float64
}

func (r *ProductRepository) GetProducts(query ProductQuery) ([]domain.Product, int64, error) {
	var products []domain.Product
	var total int64

	db := r.db.Model(&domain.Product{})

	if query.Search != "" {
		searchQuery := "%" + query.Search + "%"
		db = db.Where("name LIKE ? OR description LIKE ?", searchQuery, searchQuery)
	}

	if query.CategoryID != 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	if query.StoreID != 0 {
		db = db.Where("store_id = ?", query.StoreID)
	}
	if query.MinPrice > 0 {
		db = db.Where("price >= ?", query.MinPrice)
	}
	if query.MaxPrice > 0 {
		db = db.Where("price <= ?", query.MaxPrice)
	}

	db.Count(&total)

	if query.Limit == 0 {
		query.Limit = 10
	}
	if query.Page == 0 {
		query.Page = 1
	}

	offset := (query.Page - 1) * query.Limit
	err := db.Limit(query.Limit).Offset(offset).Find(&products).Error

	return products, total, err
}
