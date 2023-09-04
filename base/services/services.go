package services

import (
	"Society-Synergy/base/models"
)

type Service interface {
	LoginUser(*models.Login) (string, error)
	RegisterUser(*models.User) (string, error)
	GetUserByID(string) (*models.User, error)
}
