package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllSupplier() ([]models.Supplier, error) {
	suppliers := []models.Supplier{}
	data, err := s.repository.GetSupplier()
	if err != nil {
		return suppliers, err
	}
	return data, nil

}
