package services

import (
	"Society-Synergy/base/models"
)

type ServiceUser interface {

	// users
	LoginUser(*models.Login) (string, error)
	RegisterUser(*models.User) (string, models.AuditLogs, error)
	GetUserByID(string) (*models.User, error)
	SendOTP(string) error
	VerifyOTP(string, string) error
	EmailSendOTP(string) error
	VerifyOTPEmail(string, string) (string, error)
	ChangePassword(string, string, string) (models.AuditLogs, error)
	UpdateUser(string, *models.UserUpdate) (models.AuditLogs, error)

	// clubs
	GetLeaderboardByDepartment(string) (models.Club, []models.ClubMember, models.User, []models.Event, error)
	CreateDepartment(*models.CreateClub, string) (models.AuditLogs, error)

	// club members
	CreateMember(*models.CreateMember, string) (models.AuditLogs, error)

	// Sandhu
	SandhuCreateAdmin(*models.CreateAdmin, string) (models.AuditLogs, error)

	// home
	GetLeaderboardByHome() ([]models.Club, []models.Event, error)

	// event
	CreateEvent(*models.EventCreate, string) (models.Event, error)
	AddRsvp(*models.EventRSVPCreate, string) (models.EventRSVP, error)
	GetLeaderboardByEvent(string) (models.Club, models.Event, error)
}

type ServiceLogs interface {
	GetLog(*models.AuditLogs) (string, error)
	RegisterLog(*models.AuditLogs) (string, error)
}
