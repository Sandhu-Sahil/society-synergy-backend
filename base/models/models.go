package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	UserName string `json:"user_name" bson:"user_name" binding:"required,alphanum"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"userName" bson:"userName" binding:"required,alphanum"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	Password string             `json:"password" bson:"password" binding:"required"`
	Role     string             `json:"role" bson:"role"`
}
