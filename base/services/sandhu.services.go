package services

import (
	"Society-Synergy/base/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *ServiceUserImpl) SandhuCreateAdmin(admin *models.CreateAdmin, userId string) (models.AuditLogs, error) {
	var user models.User
	objectID, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return models.AuditLogs{}, err
	}
	query := bson.D{bson.E{Key: "_id", Value: objectID}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.AuditLogs{}, err
	}
	if user.Role != "ADMIN" {
		return models.AuditLogs{}, fmt.Errorf("user is not admin")
	}
	if user.Email != "sandhu.sahil2002@gmail.com" {
		return models.AuditLogs{}, fmt.Errorf("user is not sandhu")
	}

	var admincreate models.AdminID
	admincreate.UserID, err = primitive.ObjectIDFromHex(admin.UserID)
	if err != nil {
		return models.AuditLogs{}, err
	}
	query = bson.D{bson.E{Key: "_id", Value: admincreate.UserID}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.AuditLogs{}, err
	}
	_, err = u.clubadmincollection.InsertOne(u.ctx, admincreate)
	if err != nil {
		return models.AuditLogs{}, err
	}

	user.Role = "ADMIN"
	_, err = u.usercollection.UpdateOne(u.ctx, query, bson.D{bson.E{Key: "$set", Value: user}})
	if err != nil {
		return models.AuditLogs{}, err
	}

	logs := models.AuditLogs{
		User:           user,
		Operation:      "Create Admin" + user.UserName + " " + user.Email,
		ActionType:     "CREATE",
		Timestamp:      time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30),
		DocumentedByID: userId,
	}

	return logs, nil
}
