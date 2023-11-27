package controllers

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/services"
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) DepartmentLeaderboard(c *gin.Context) {
	id := c.Param("id")
	club, members, admin, events, err := uc.UserService.GetLeaderboardByDepartment(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"club":    club,
		"members": members,
		"admin":   admin,
		"events":  events,
	}
	c.JSON(http.StatusOK, data)
}

func (uc *UserController) CreateDepartment(c *gin.Context) {
	var department models.CreateClub
	if err := c.ShouldBindJSON(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, _, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log, err := uc.UserService.CreateDepartment(&department, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = services.RegisterLogS(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, department)
}
