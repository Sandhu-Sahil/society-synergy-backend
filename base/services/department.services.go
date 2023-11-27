package services

import (
	"Society-Synergy/base/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *ServiceUserImpl) GetLeaderboardByDepartment(id string) (models.Club, []models.ClubMember, models.User, []models.Event, error) {
	var club models.Club
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}

	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	err = u.clubcollection.FindOne(u.ctx, query).Decode(&club)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}

	var members []models.ClubMember
	query = bson.D{bson.E{Key: "clubID", Value: objectid}}
	cursor, err := u.clubmembercollection.Find(u.ctx, query)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}
	if err = cursor.All(u.ctx, &members); err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}

	var adminid models.AdminID
	query = bson.D{bson.E{Key: "_id", Value: club.AdminID}}
	err = u.clubadmincollection.FindOne(u.ctx, query).Decode(&adminid)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}

	var admin models.User
	query = bson.D{bson.E{Key: "_id", Value: adminid.UserID}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&admin)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}
	admin.Password = ""
	admin.OTP = ""
	admin.OTPExpiry = time.Time{}
	admin.PhoneNo = ""

	query = bson.D{bson.E{Key: "clubID", Value: objectid}}
	options := options.Find()
	options.SetSort(bson.D{{"startDate", -1}})
	options.SetLimit(15)

	cursor, err = u.eventcollection.Find(u.ctx, query, options)
	if err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}
	var events []models.Event
	if err = cursor.All(u.ctx, &events); err != nil {
		return models.Club{}, nil, models.User{}, nil, err
	}

	return club, members, admin, events, nil
}

func (u *ServiceUserImpl) CreateDepartment(department *models.CreateClub, user_id string) (models.AuditLogs, error) {
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

	var club models.Club
	club.AdminID, err = primitive.ObjectIDFromHex(department.AdminID)
	if err != nil {
		return models.AuditLogs{}, err
	}

	// check if admin exists
	var admin models.AdminID
	query = bson.D{bson.E{Key: "_id", Value: club.AdminID}}
	err = u.clubadmincollection.FindOne(u.ctx, query).Decode(&admin)
	if err != nil {
		return models.AuditLogs{}, err
	}
	// check if department already exists
	query = bson.D{bson.E{Key: "name", Value: department.Name}}
	err = u.clubcollection.FindOne(u.ctx, query).Decode(&club)
	if err == nil {
		return models.AuditLogs{}, fmt.Errorf("department already exists")
	}

	club.Description = department.Description
	club.Name = department.Name
	club.Github = department.Github
	club.Instagram = department.Instagram
	club.LinkedIn = department.LinkedIn
	club.Website = department.Website
	club.LogoUrl = department.LogoUrl

	_, err = u.clubcollection.InsertOne(u.ctx, club)
	if err != nil {
		return models.AuditLogs{}, err
	}

	log := models.AuditLogs{
		User:           user,
		ActionType:     "CREATE",
		Operation:      "Created a new department by the name " + department.Name + " with description " + department.Description + ".",
		Timestamp:      time.Now().UTC().Add(time.Hour * 5).Add(time.Minute * 30),
		DocumentedByID: user_id,
	}

	return log, nil
}
