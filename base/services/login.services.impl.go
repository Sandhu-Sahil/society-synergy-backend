package services

import (
	"errors"
	"strings"

	"Society-Synergy/base/models"
	"Society-Synergy/base/token"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *ServiceImpl) RegisterUser(user *models.User) (string, error) {
	query := bson.D{bson.E{Key: "user_name", Value: user.UserName}}
	res, err := u.usercollection.Find(u.ctx, query)
	if err != nil {
		return "", err
	}
	// fmt.Print(res.RemainingBatchLength())
	if res.RemainingBatchLength() != 0 {
		return "", errors.New("user already existed (user name already registered)")
	}
	query = bson.D{bson.E{Key: "email", Value: user.Email}}
	res, err = u.usercollection.Find(u.ctx, query)
	if err != nil {
		return "", err
	}
	// fmt.Print(res.RemainingBatchLength())
	if res.RemainingBatchLength() != 0 {
		return "", errors.New("user already existed (email already registered)")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.UserName = strings.TrimSpace(user.UserName)

	userCreated, err := u.usercollection.InsertOne(u.ctx, user)
	if err != nil {
		return "", err
	}
	id := userCreated.InsertedID.(primitive.ObjectID).Hex()

	token, err := token.GenerateToken(id, user.UserName, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *ServiceImpl) LoginUser(user *models.Login) (string, error) {
	var userFound *models.User
	query := bson.D{bson.E{Key: "userName", Value: user.UserName}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}

	token, err := token.GenerateToken(userFound.ID.Hex(), user.UserName, userFound.Role)
	if err != nil {
		return "", err
	}

	return token, err
}
