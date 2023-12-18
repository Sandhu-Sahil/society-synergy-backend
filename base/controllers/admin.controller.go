package controllers

import (
	"Society-Synergy/base/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) AdminLogin(ctx *gin.Context) {
	var user models.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.UserService.AdminLogin(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := map[string]interface{}{
		"token": token,
		// "isteam": team,
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Login verified", "data": data})
}
