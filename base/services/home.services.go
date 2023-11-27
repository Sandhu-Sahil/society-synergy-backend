package services

import (
	"Society-Synergy/base/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (u *ServiceUserImpl) GetLeaderboardByHome() ([]models.Club, []models.Event, error) {
	var clubs []models.Club
	// extract all clubs
	cursor, err := u.clubcollection.Find(u.ctx, bson.M{})
	if err != nil {
		return nil, nil, err
	}
	if err = cursor.All(u.ctx, &clubs); err != nil {
		return nil, nil, err
	}

	// extract top 15 recent events added in database
	query := bson.D{{}}
	options := options.Find()
	options.SetSort(bson.D{{"startDate", -1}})
	options.SetLimit(15)

	cursor, err = u.eventcollection.Find(u.ctx, query, options)
	if err != nil {
		return nil, nil, err
	}
	var events []models.Event
	if err = cursor.All(u.ctx, &events); err != nil {
		return nil, nil, err
	}

	return clubs, events, nil
}
