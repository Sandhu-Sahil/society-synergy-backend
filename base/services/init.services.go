package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceUserImpl struct {
	usercollection       *mongo.Collection
	clubcollection       *mongo.Collection
	clubmembercollection *mongo.Collection
	clubadmincollection  *mongo.Collection
	ctx                  context.Context
}

type ServiceLogsImpl struct {
	logcollection *mongo.Collection
	ctx           context.Context
}

func NewServiceUser(usercollection *mongo.Collection, clubcollection *mongo.Collection, clubmembercollection *mongo.Collection, clubadmincollection *mongo.Collection, ctx context.Context) ServiceUser {
	return &ServiceUserImpl{
		usercollection:       usercollection,
		clubcollection:       clubcollection,
		clubmembercollection: clubmembercollection,
		clubadmincollection:  clubadmincollection,
		ctx:                  ctx,
	}
}

func NewServiceLogs(logcollection *mongo.Collection, ctx context.Context) ServiceLogs {
	LogCollection = logcollection
	Ctx = ctx
	return &ServiceLogsImpl{
		logcollection: logcollection,
		ctx:           ctx,
	}
}
