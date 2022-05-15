package repository

import "MINI_PROJECT_ALTERRA/models"

func (r *Repository) GetSupplier() ([]models.Supplier, error) {
	suppliers := []models.Supplier{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&suppliers).Error
	if err != nil {
		return suppliers, err
	}
	return suppliers, nil

}
