package controller

import "MINI_PROJECT_ALTERRA/service"

type AllController struct {
	service service.Service
}

func NewAllController(service service.Service) *AllController {
	return &AllController{service: service}
}
