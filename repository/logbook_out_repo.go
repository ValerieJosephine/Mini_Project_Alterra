package repository

import "MINI_PROJECT_ALTERRA/models"

func (r *Repository) GetStockOut() ([]models.StockOut, error) {
	logbook_stock_out := []models.StockOut{}
	// err := r.db.Raw("select * from users").Scan(&users).Error
	err := r.db.Find(&logbook_stock_out).Error
	if err != nil {
		return logbook_stock_out, err
	}
	return logbook_stock_out, nil

}
