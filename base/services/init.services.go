package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

func NewService(usercollection *mongo.Collection, ctx context.Context) Service {
	return &ServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
