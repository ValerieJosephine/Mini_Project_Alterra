package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllProductType() ([]models.ProductType, error) {
	product_types := []models.ProductType{}
	data, err := s.repository.GetProductType()
	if err != nil {
		return product_types, err
	}
	return data, nil

}
