package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	UserName string `json:"userName" bson:"userName" binding:"required,alphanum"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName string             `json:"userName" bson:"userName" binding:"required,alphanum"`
	Email    string             `json:"email" bson:"email" binding:"required,email"`
	Password string             `json:"password" bson:"password" binding:"required"`
	Role     string             `json:"role" bson:"role"`
	// CountryCode string             `json:"countryCode" bson:"countryCode" binding:"required,iso3166_1_alpha2"`
	PhoneNo   string `json:"phoneNo" bson:"phoneNo" binding:"required,e164,min=13,max=13"`
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required"`
}

type Club struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	AdminID     User               `json:"adminID" bson:"adminID" binding:"required"`
}

type ClubMember struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClubID Club               `json:"clubID" bson:"clubID" binding:"required"`
	UserID User               `json:"userID" bson:"userID" binding:"required"`
}

type Event struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClubID       Club               `json:"clubID" bson:"clubID" binding:"required"`
	Name         string             `json:"name" bson:"name" binding:"required"`
	Description  string             `json:"description" bson:"description" binding:"required"`
	StartDate    string             `json:"startDate" bson:"startDate" binding:"required"`
	EndDate      string             `json:"endDate" bson:"endDate" binding:"required"`
	EventTime    string             `json:"eventTime" bson:"eventTime" binding:"required"`
	Location     string             `json:"location" bson:"location" binding:"required"`
	RSVPDeadline string             `json:"rsvpDeadline" bson:"rsvpDeadline" binding:"required"`
}

type EventRSVP struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	EventID Event              `json:"eventID" bson:"eventID" binding:"required"`
	UserID  User               `json:"userID" bson:"userID" binding:"required"`
	Status  string             `json:"status" bson:"status" binding:"required"`
}

// type EventRSVPCount struct {
// 	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Event  Event              `json:"event" bson:"event" binding:"required"`
// 	Count  int                `json:"count" bson:"count" binding:"required"`
// 	Status string             `json:"status" bson:"status" binding:"required"`
// }

type AdminID struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserID User               `json:"userID" bson:"userID" binding:"required"`
}

type Email struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Email         string             `json:"email" bson:"email" binding:"required,email"`
	SenderAdminID AdminID            `json:"senderAdminID" bson:"senderAdminID" binding:"required"`
	RecipientIDs  []User             `json:"recipientIDs" bson:"recipientIDs" binding:"required"`
	Subject       string             `json:"subject" bson:"subject" binding:"required"`
	Message       string             `json:"message" bson:"message" binding:"required"`
	Timestamp     string             `json:"timestamp" bson:"timestamp" binding:"required"`
}

type AuditLogs struct {
	// baad ch karange
}
