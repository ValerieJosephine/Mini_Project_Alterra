package service

import "MINI_PROJECT_ALTERRA/repository"

type Service struct {
	repository repository.AllRepository
}

func NewAllService(repository repository.AllRepository) *Service {
	return &Service{repository: repository}
}
