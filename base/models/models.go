package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Login struct {
	Email    string `json:"email" bson:"email" binding:"required,email"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	UserName  string             `json:"userName" bson:"userName" binding:"required"`
	Email     string             `json:"email" bson:"email" binding:"required,email"`
	Password  string             `json:"password" bson:"password" binding:"required"`
	Role      string             `json:"role" bson:"role"` // student, head, admin
	PhoneNo   string             `json:"phoneNo" bson:"phoneNo" binding:"required,e164,min=13,max=13"`
	FirstName string             `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string             `json:"lastName" bson:"lastName" binding:"required"`
	Varified  bool               `json:"varified" bson:"varified"`
	OTP       string             `json:"otp" bson:"otp"`
	OTPExpiry time.Time          `json:"otpExpiry" bson:"otpExpiry"`
}

type Club struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name" binding:"required"`
	Description string             `json:"description" bson:"description" binding:"required"`
	AdminID     primitive.ObjectID `json:"adminID" bson:"adminID" binding:"required"`
	Instagram   string             `json:"instagram" bson:"instagram"`
	LinkedIn    string             `json:"linkedIn" bson:"linkedIn"`
	Github      string             `json:"github" bson:"github"`
	Website     string             `json:"website" bson:"website"`
}

type ClubMember struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClubID    primitive.ObjectID `json:"clubID" bson:"clubID" binding:"required"`
	UserID    primitive.ObjectID `json:"userID" bson:"userID" binding:"required"`
	Name      string             `json:"name" bson:"name" binding:"required"`
	Instagram string             `json:"instagram" bson:"instagram" binding:"required"`
	LinkedIn  string             `json:"linkedIn" bson:"linkedIn" binding:"required"`
	Github    string             `json:"github" bson:"github" binding:"required"`
	Role      string             `json:"role" bson:"role" binding:"required"`
}

type Event struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ClubID       primitive.ObjectID `json:"clubID" bson:"clubID" binding:"required"`
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
	EventID primitive.ObjectID `json:"eventID" bson:"eventID" binding:"required"`
	UserID  primitive.ObjectID `json:"userID" bson:"userID" binding:"required"`
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
	UserID primitive.ObjectID `json:"userID" bson:"userID" binding:"required"`
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
	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User           User               `json:"user" bson:"user" binding:"required"`
	ActionType     string             `json:"actionType" bson:"actionType" binding:"required"` // delete, update, create
	Operation      string             `json:"operation" bson:"operation" binding:"required"`
	Timestamp      time.Time          `json:"timestamp" bson:"timestamp" binding:"required"`
	DocumentedByID string             `json:"documentID" bson:"documentID" binding:"required"`
	BeforeEdit     interface{}        `json:"beforeEdit" bson:"beforeEdit"`
	AfterEdit      interface{}        `json:"afterEdit" bson:"afterEdit"`
}

type UserUpdate struct {
	Otp       string `json:"otp"`
	UserName  string `json:"userName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	PhoneNo   string `json:"phoneNo"`
}

type CreateClub struct {
	Name        string `json:"name" bson:"name" binding:"required"`
	Description string `json:"description" bson:"description" binding:"required"`
	AdminID     string `json:"adminID" bson:"adminID" binding:"required"`
	Instagram   string `json:"instagram" bson:"instagram" binding:"required"`
	LinkedIn    string `json:"linkedIn" bson:"linkedIn"`
	Github      string `json:"github" bson:"github"`
	Website     string `json:"website" bson:"website"`
}

type CreateAdmin struct {
	UserID string `json:"userID" bson:"userID" binding:"required"`
}

type CreateMember struct {
	ClubID    string `json:"clubID" bson:"clubID" binding:"required"`
	UserID    string `json:"userID" bson:"userID" binding:"required"`
	Instagram string `json:"instagram" bson:"instagram" binding:"required"`
	LinkedIn  string `json:"linkedIn" bson:"linkedIn" binding:"required"`
	Github    string `json:"github" bson:"github" binding:"required"`
	Role      string `json:"role" bson:"role" binding:"required"`
}
