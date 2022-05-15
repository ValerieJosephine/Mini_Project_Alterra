package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllStockIn() ([]models.StockIn, error) {
	logbook_stock_in := []models.StockIn{}
	data, err := s.repository.GetStockIn()
	if err != nil {
		return logbook_stock_in, err
	}
	return data, nil

}
