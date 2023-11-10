package services

import (
	"Society-Synergy/base/models"

	"go.mongodb.org/mongo-driver/bson"
)

func (u *ServiceUserImpl) GetLeaderboardByHome() ([]models.Club, error) {
	var clubs []models.Club
	// extract all clubs
	cursor, err := u.clubcollection.Find(u.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(u.ctx, &clubs); err != nil {
		return nil, err
	}
	return clubs, nil
}
