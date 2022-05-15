package repository

import "MINI_PROJECT_ALTERRA/models"

func (r *Repository) GetProductType() ([]models.ProductType, error) {
	product_types := []models.ProductType{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&product_types).Error
	if err != nil {
		return product_types, err
	}
	return product_types, nil

}
