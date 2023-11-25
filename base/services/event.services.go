package services

import (
	"Society-Synergy/base/models"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (u *ServiceUserImpl) CreateEvent(event *models.EventCreate, user_id string) (models.Event, error) {
	objectid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return models.Event{}, err
	}

	var user models.User
	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.Event{}, err
	}
	if user.Role == "STUDENT" {
		return models.Event{}, fmt.Errorf("user is not admin or head")
	}

	// check if club exists
	var club models.Club
	objectid1, err := primitive.ObjectIDFromHex(event.ClubID)
	if err != nil {
		return models.Event{}, err
	}
	query = bson.D{bson.E{Key: "_id", Value: objectid1}}
	err = u.clubcollection.FindOne(u.ctx, query).Decode(&club)
	if err != nil {
		return models.Event{}, err
	}

	// create event
	newevent := models.Event{
		ClubID:       objectid1,
		Name:         event.Name,
		Description:  event.Description,
		StartDate:    event.StartDate,
		EndDate:      event.EndDate,
		EventTime:    event.EventTime,
		Location:     event.Location,
		RSVPDeadline: event.RSVPDeadline,
		PosterUrl:    event.PosterUrl,
	}
	_, err = u.eventcollection.InsertOne(u.ctx, newevent)
	if err != nil {
		return models.Event{}, err
	}

	return newevent, nil
}

func (u *ServiceUserImpl) AddRsvp(rsvp *models.EventRSVPCreate, user_id string) (models.EventRSVP, error) {
	objectid, err := primitive.ObjectIDFromHex(user_id)
	if err != nil {
		return models.EventRSVP{}, err
	}

	var user models.User
	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	err = u.usercollection.FindOne(u.ctx, query).Decode(&user)
	if err != nil {
		return models.EventRSVP{}, err
	}

	// check if event exists
	var event models.Event
	objectid1, err := primitive.ObjectIDFromHex(rsvp.EventID)
	if err != nil {
		return models.EventRSVP{}, err
	}
	query = bson.D{bson.E{Key: "_id", Value: objectid1}}
	err = u.eventcollection.FindOne(u.ctx, query).Decode(&event)
	if err != nil {
		return models.EventRSVP{}, err
	}

	// check if user already rsvped
	var eventrsvp models.EventRSVP
	query = bson.D{bson.E{Key: "eventID", Value: objectid1}, bson.E{Key: "userID", Value: objectid}}
	err = u.eventrsvpcollection.FindOne(u.ctx, query).Decode(&eventrsvp)
	if err == nil {
		return models.EventRSVP{}, fmt.Errorf("user already rsvped")
	}

	// check date of rsvp
	// convert string to time and check if rsvp deadline has passed
	layout := "2006-01-02T15:04:05.000Z"
	t, err := time.Parse(layout, event.RSVPDeadline)
	if err != nil {
		return models.EventRSVP{}, err
	}
	if time.Now().After(t) {
		return models.EventRSVP{}, fmt.Errorf("rsvp deadline has passed")
	}

	// create event rsvp
	neweventrsvp := models.EventRSVP{
		EventID:      objectid1,
		EventDetails: event,
		UserID:       objectid,
	}
	_, err = u.eventrsvpcollection.InsertOne(u.ctx, neweventrsvp)
	if err != nil {
		return models.EventRSVP{}, err
	}

	return neweventrsvp, nil
}
