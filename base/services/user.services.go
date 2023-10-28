package services

import (
	"Society-Synergy/base/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *ServiceUserImpl) GetUserByID(id string) (*models.User, error) {
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
	userFound.OTP = "**PROTECTED**"
	return userFound, nil
}

func (u *ServiceUserImpl) ChangePassword(user_id string, otp string, newPassword string) error {
	objectid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err = u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return err
	}

	if userFound.OTP != otp {
		return fmt.Errorf("invalid OTP")
	}

	if time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30).After(userFound.OTPExpiry) {
		return fmt.Errorf("OTP expired")
	}

	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userFound.Password = string(hashedPassword)

	_, err = u.usercollection.UpdateOne(u.ctx, query, bson.D{bson.E{Key: "$set", Value: userFound}})
	if err != nil {
		return err
	}

	return nil
}

func (u *ServiceUserImpl) UpdateUser(user_id string, update *models.UserUpdate) error {
	objectid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	var userFound *models.User
	err = u.usercollection.FindOne(u.ctx, query).Decode(&userFound)
	if err != nil {
		return err
	}

	if userFound.OTP != update.Otp {
		return fmt.Errorf("invalid OTP")
	}

	if time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30).After(userFound.OTPExpiry) {
		return fmt.Errorf("OTP expired")
	}

	// update user
	userFound.FirstName = update.FirstName
	userFound.LastName = update.LastName
	userFound.PhoneNo = update.PhoneNo
	userFound.UserName = update.UserName

	_, err = u.usercollection.UpdateOne(u.ctx, query, bson.D{bson.E{Key: "$set", Value: userFound}})
	if err != nil {
		return err
	}

	return nil
}
