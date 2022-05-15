package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllTypeSupplier() ([]models.SupplierType, error) {
	supplier_category := []models.SupplierType{}
	data, err := s.repository.GetSupplierType()
	if err != nil {
		return supplier_category, err
	}
	return data, nil

}
