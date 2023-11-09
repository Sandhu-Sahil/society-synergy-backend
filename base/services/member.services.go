package services

import (
	"Society-Synergy/base/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *ServiceUserImpl) CreateMember(member *models.CreateMember, user_id string) (models.AuditLogs, error) {
	objectid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return models.AuditLogs{}, err
	}

	var user models.User
	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.AuditLogs{}, err
	}
	if user.Role != "ADMIN" {
		return models.AuditLogs{}, fmt.Errorf("user is not admin")
	}

	// check if club exists
	var club models.Club
	objectid1, err := primitive.ObjectIDFromHex(member.ClubID)
	if err != nil {
		return models.AuditLogs{}, err
	}
	query = bson.D{bson.E{Key: "_id", Value: objectid1}}
	err = u.clubcollection.FindOne(u.ctx, query).Decode(&club)
	if err != nil {
		return models.AuditLogs{}, err
	}

	// check if user exists
	objectid2, err := primitive.ObjectIDFromHex(member.UserID)
	if err != nil {
		return models.AuditLogs{}, err
	}
	query = bson.D{bson.E{Key: "_id", Value: objectid2}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.AuditLogs{}, err
	}

	// check if user is already a member
	var clubmember models.ClubMember
	query = bson.D{bson.E{Key: "clubID", Value: objectid1}, bson.E{Key: "userID", Value: objectid2}}
	err = u.clubmembercollection.FindOne(u.ctx, query).Decode(&clubmember)
	if err == nil {
		return models.AuditLogs{}, fmt.Errorf("user is already a member")
	}

	// create club member
	clubmember = models.ClubMember{
		ClubID:    objectid1,
		UserID:    objectid2,
		Name:      user.FirstName + " " + user.LastName,
		Instagram: member.Instagram,
		LinkedIn:  member.LinkedIn,
		Github:    member.Github,
		Role:      member.Role,
	}
	_, err = u.clubmembercollection.InsertOne(u.ctx, clubmember)
	if err != nil {
		return models.AuditLogs{}, err
	}

	// update user role
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "role", Value: "HEAD"}}}}
	query = bson.D{bson.E{Key: "_id", Value: objectid2}}
	_, err = u.usercollection.UpdateOne(u.ctx, query, update)
	if err != nil {
		return models.AuditLogs{}, err
	}

	// create log
	log := models.AuditLogs{
		User:           user,
		ActionType:     "CREATE",
		Operation:      "Member " + user.UserName + " added to club " + club.Name,
		Timestamp:      time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30),
		DocumentedByID: user_id,
	}

	return log, nil
}
