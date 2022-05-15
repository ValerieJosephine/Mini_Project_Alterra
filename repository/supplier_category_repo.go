package repository

import "MINI_PROJECT_ALTERRA/models"

func (r *Repository) GetSupplierType() ([]models.SupplierType, error) {
	supplier_category := []models.SupplierType{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&supplier_category).Error
	if err != nil {
		return supplier_category, err
	}
	return supplier_category, nil

}
