package service

import (
	"MINI_PROJECT_ALTERRA/models"
	"MINI_PROJECT_ALTERRA/repository"
)

type Service struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *Service {
	return &Service{repository: repository}
}

func (s *Service) FetchAllproduct() ([]models.Product, error) {
	var products []models.Product
	data, err := s.repository.FetchProducts()
	if err != nil {
		return products, err
	}

	return data, nil

}

func (s *Service) StoreAllProduct(input models.Product) (models.Product, error) {
	data, err := s.repository.StoreProduct(input)
	if err != nil {
		return input, err
	}

	return data, nil

}
