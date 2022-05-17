package service

import (
	"MINI_PROJECT_ALTERRA/models"
	// "MINI_PROJECT_ALTERRA/repository"
)

// type Service struct {
// 	repository repository.AllRepository
// }

// func NewAllService(repository repository.AllRepository) *Service {
// 	return &Service{repository: repository}
// }

func (s *Service) FetchAllproduct() ([]models.Product, error) {
	var products []models.Product
	data, err := s.repository.FetchProducts()
	if err != nil {
		return products, err
	}

	return data, nil

}

func (s *Service) FetchProductWithId(id int) (models.Product, error) {
	data, err := s.repository.FetchProductById(id)
	if err != nil {
		return data, err
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

func (s *Service) DeleteAllProduct(id int) error {
	err := s.repository.DeleteProduct(id)
	if err != nil {
		return err
	}

	return err
}

// func (s *Service) UpdateProducts() ([]models.Product, error) {
// 	var products []models.Product
// 	data, err := s.repository.UpdateProduct(&products)
// 	if err != nil {
// 		return products, err
// 	}

// 	return data, nil

// }
