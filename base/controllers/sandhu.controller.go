package controllers

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/services"
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) SandhuCreateAdmin(c *gin.Context) {
	var admin models.CreateAdmin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, _, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	log, err := uc.UserService.SandhuCreateAdmin(&admin, user_id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.RegisterLogS(&log)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admin)
}
