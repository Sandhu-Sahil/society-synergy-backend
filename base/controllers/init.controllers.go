package controllers

import "Society-Synergy/base/services"

type Controller struct {
	UserService services.Service
}

func New(userservice services.Service) Controller {
	return Controller{
		UserService: userservice,
	}
}
