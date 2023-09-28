package controllers

import "Society-Synergy/base/services"

type UserController struct {
	UserService services.ServiceUser
}

type LogsController struct {
	LogsService services.ServiceLogs
}

func NewUserController(userservice services.ServiceUser) UserController {
	return UserController{
		UserService: userservice,
	}
}

func NewLogsController(logsservice services.ServiceLogs) LogsController {
	return LogsController{
		LogsService: logsservice,
	}
}
