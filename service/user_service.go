package service

import "MINI_PROJECT_ALTERRA/models"

func (s *Service) GetAllUser() ([]models.User, error) {
	users := []models.User{}
	data, err := s.repository.GetUsers()
	if err != nil {
		return users, err
	}
	return data, nil

}
