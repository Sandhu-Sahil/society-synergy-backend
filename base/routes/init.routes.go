package routes

import "Society-Synergy/base/controllers"

type RouterService struct {
	UserController controllers.UserController
	LogsController controllers.LogsController
}

func NewRouterService(UserController controllers.UserController, LogsController controllers.LogsController) RouterService {
	return RouterService{
		UserController: UserController,
		LogsController: LogsController,
	}
}
