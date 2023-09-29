package services

import (
	"Society-Synergy/base/models"
)

type ServiceUser interface {
	LoginUser(*models.Login) (string, error)
	RegisterUser(*models.User) (string, models.AuditLogs, error)
	GetUserByID(string) (*models.User, error)
	SendOTP(string) error
	VerifyOTP(string, string) error
}

type ServiceLogs interface {
	GetLog(*models.AuditLogs) (string, error)
	RegisterLog(*models.AuditLogs) (string, error)
}
