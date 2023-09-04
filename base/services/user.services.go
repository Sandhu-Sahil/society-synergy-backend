package services

import (
	"Society-Synergy/base/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *ServiceImpl) GetUserByID(id string) (*models.User, error) {
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &models.User{}, err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err = u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return &models.User{}, err
	}

	userFound.Password = "**PROTECTED**"
	return userFound, nil
}
