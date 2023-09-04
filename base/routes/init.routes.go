package routes

import "Society-Synergy/base/controllers"

type RouterService struct {
	Controller controllers.Controller
}

func NewRouterService(Controller controllers.Controller) RouterService {
	return RouterService{
		Controller: Controller,
	}
}
