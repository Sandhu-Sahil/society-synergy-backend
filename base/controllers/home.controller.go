package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) HomeLeaderboard(c *gin.Context) {
	clubs, events, err := uc.UserService.GetLeaderboardByHome()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]interface{}{
		"clubs":  clubs,
		"events": events,
	}
	c.JSON(http.StatusOK, data)
}
