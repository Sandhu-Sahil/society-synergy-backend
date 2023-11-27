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
	// conevrt string to time
	layout := "2006-01-02"
	t, err := time.Parse(layout, event.StartDate)
	if err != nil {
		return models.Event{}, err
	}
	// check if start date is after today
	if time.Now().After(t) {
		return models.Event{}, fmt.Errorf("start date is before today")
	}
	// do same with end date and rsvp deadline
	t, err = time.Parse(layout, event.EndDate)
	if err != nil {
		return models.Event{}, err
	}
	if time.Now().After(t) {
		return models.Event{}, fmt.Errorf("end date is before today")
	}
	t, err = time.Parse(layout, event.RSVPDeadline)
	if err != nil {
		return models.Event{}, err
	}
	if time.Now().After(t) {
		return models.Event{}, fmt.Errorf("rsvp deadline is before today")
	}

	if event.PosterUrl == "" {
		event.PosterUrl = "https://cdn.discordapp.com/attachments/888075387228278817/1178710978217639966/Event.png"
	}
	// check if poster url is valid
	if err := urlValidator(event.PosterUrl); err != nil {
		return models.Event{}, err
	}

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

func urlValidator(url string) error {
	// check if url is valid
	// check if url end with .png, .jpg, .jpeg, .gif, .webp
	if len(url) < 4 {
		return fmt.Errorf("invalid url")
	}
	if url[len(url)-4:] == ".png" || url[len(url)-4:] == ".jpg" || url[len(url)-4:] == ".gif" || url[len(url)-4:] == ".webp" {
		return nil
	}
	if url[len(url)-5:] == ".jpeg" {
		return nil
	}
	return fmt.Errorf("invalid url")
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
		return models.EventRSVP{}, fmt.Errorf("user already rsvped, check your rsvp list on your profile")
	}

	// check date of rsvp
	// convert string to time and check if rsvp deadline has passed
	layout := "2006-01-02"
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

func (u *ServiceUserImpl) GetLeaderboardByEvent(id string) (models.Club, models.Event, error) {
	objectid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Club{}, models.Event{}, err
	}

	var event models.Event
	query := bson.D{bson.E{Key: "_id", Value: objectid}}
	err = u.eventcollection.FindOne(u.ctx, query).Decode(&event)
	if err != nil {
		return models.Club{}, models.Event{}, err
	}

	var club models.Club
	query = bson.D{bson.E{Key: "_id", Value: event.ClubID}}
	err = u.clubcollection.FindOne(u.ctx, query).Decode(&club)
	if err != nil {
		return models.Club{}, models.Event{}, err
	}

	return club, event, nil
}
