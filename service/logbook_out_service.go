package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllStockOut() ([]models.StockOut, error) {
	logbook_stock_out := []models.StockOut{}
	data, err := s.repository.GetStockOut()
	if err != nil {
		return logbook_stock_out, err
	}
	return data, nil

}
