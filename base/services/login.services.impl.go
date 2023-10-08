package services

import (
	"errors"
	"strings"
	"time"

	"Society-Synergy/base/models"
	"Society-Synergy/base/token"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func (u *ServiceUserImpl) RegisterUser(user *models.User) (string, models.AuditLogs, error) {
	query := bson.D{bson.E{Key: "userName", Value: user.UserName}}
	res, err := u.usercollection.Find(u.ctx, query)
	if err != nil {
		return "", models.AuditLogs{}, err
	}
	// fmt.Print(res.RemainingBatchLength())
	if res.RemainingBatchLength() != 0 {
		return "", models.AuditLogs{}, errors.New("user already existed (user name already registered)")
	}
	query = bson.D{bson.E{Key: "email", Value: user.Email}}
	res, err = u.usercollection.Find(u.ctx, query)
	if err != nil {
		return "", models.AuditLogs{}, err
	}
	// fmt.Print(res.RemainingBatchLength())
	if res.RemainingBatchLength() != 0 {
		return "", models.AuditLogs{}, errors.New("user already existed (email already registered)")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", models.AuditLogs{}, err
	}
	user.Password = string(hashedPassword)

	//remove spaces in username
	user.UserName = strings.TrimSpace(user.UserName)
	user.OTP = ""
	user.OTPExpiry = time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30)
	user.Varified = false

	userCreated, err := u.usercollection.InsertOne(u.ctx, user)
	if err != nil {
		return "", models.AuditLogs{}, err
	}
	id := userCreated.InsertedID.(primitive.ObjectID).Hex()

	token, err := token.GenerateToken(id, user.UserName, user.Role)
	if err != nil {
		return "", models.AuditLogs{}, err
	}

	log := models.AuditLogs{
		User:           *user,
		ActionType:     "CREATE",
		Operation:      "Created a new user by the name " + user.UserName + " with email " + user.Email + " and role " + user.Role + ".",
		Timestamp:      time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30),
		DocumentedByID: id,
	}

	return token, log, nil
}

func (u *ServiceUserImpl) LoginUser(user *models.Login) (string, error) {
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

	token, err := token.GenerateToken(userFound.ID.Hex(), userFound.UserName, userFound.Role)
	if err != nil {
		return "", err
	}

	return token, err
}
