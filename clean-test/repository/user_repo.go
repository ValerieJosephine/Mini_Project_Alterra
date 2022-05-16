package repository

import (
	"MINI_PROJECT_ALTERRA/models"

	"gorm.io/gorm"
)

type AllRepository interface {
	FetchProducts() ([]models.Product, error)
	StoreProduct(input models.Product) (models.Product, error)
	DeleteProduct(id int) error
	GetProductType() ([]models.ProductType, error)
	GetSupplier() ([]models.Supplier, error)
	GetSupplierType() ([]models.SupplierType, error)
	GetStockIn() ([]models.StockIn, error)
	GetStockOut() ([]models.StockOut, error)
	GetUsers() ([]models.User, error)
}

type Repository struct {
	db *gorm.DB
}

func NewAllRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}

}
func (r *Repository) GetUsers() ([]models.User, error) {
	users := []models.User{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&users).Error
	if err != nil {
		return users, err
	}
	return users, nil

}
