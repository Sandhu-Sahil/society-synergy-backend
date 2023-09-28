package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceUserImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

type ServiceLogsImpl struct {
	logcollection *mongo.Collection
	ctx           context.Context
}

func NewServiceUser(usercollection *mongo.Collection, ctx context.Context) ServiceUser {
	return &ServiceUserImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}

func NewServiceLogs(logcollection *mongo.Collection, ctx context.Context) ServiceLogs {
	return &ServiceLogsImpl{
		logcollection: logcollection,
		ctx:           ctx,
	}
}
