package services

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/token"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func (u *ServiceUserImpl) AdminLogin(user *models.Login) (string, error) {
	var userFound *models.User
	query := bson.D{bson.E{Key: "email", Value: user.Email}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	// check role and return new error
	if userFound.Role != "ADMIN" && userFound.Role != "HEAD" {
		return "", errors.New("user is not admin or head")
	}

	token, err := token.GenerateToken(userFound.ID.Hex(), userFound.UserName, userFound.Role)
	if err != nil {
		return "", err
	}

	return token, err
}
