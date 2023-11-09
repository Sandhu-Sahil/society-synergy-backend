package controllers

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/services"
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) CreateMember(c *gin.Context) {
	var member models.CreateMember
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user_id, _, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log, err := uc.UserService.CreateMember(&member, user_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err = services.RegisterLogS(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, member)
}
