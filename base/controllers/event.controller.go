package controllers

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (us *UserController) CreateEvent(c *gin.Context) {
	var event models.EventCreate
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if event.Name == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, name is missing"})
		return
	}
	if event.Description == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, description is missing"})
		return
	}
	if event.StartDate == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, startDate is missing"})
		return
	}
	if event.EndDate == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, endDate is missing"})
		return
	}
	if event.EventTime == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, eventTime is missing"})
		return
	}
	if event.Location == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, location is missing"})
		return
	}
	if event.RSVPDeadline == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, rsvpDeadline is missing"})
		return
	}
	if event.PosterUrl == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, posterUrl is missing"})
		return
	}

	user_id, _, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newevent, err := us.UserService.CreateEvent(&event, user_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Event created successfully", "data": newevent})
}

func (us *UserController) AddRsvp(c *gin.Context) {
	var event models.EventRSVPCreate
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if event.EventID == "" {
		c.JSON(400, gin.H{"error": "must provide all fields, eventID is missing"})
		return
	}

	user_id, _, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newevent, err := us.UserService.AddRsvp(&event, user_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Rsvp added successfully", "data": newevent})
}
