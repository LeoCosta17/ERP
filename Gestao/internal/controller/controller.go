package controller

import "gestao/internal/service"

type Controller struct {
}

func NewController(service *service.Service) *Controller {
	return &Controller{}
}
