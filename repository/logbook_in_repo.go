package repository

import "MINI_PROJECT_ALTERRA/models"

func (r *Repository) GetStockIn() ([]models.StockIn, error) {
	logbook_stock_in := []models.StockIn{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&logbook_stock_in).Error
	if err != nil {
		return logbook_stock_in, err
	}
	return logbook_stock_in, nil

}
