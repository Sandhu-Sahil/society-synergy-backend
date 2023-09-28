package controllers

import (
	"errors"
	"net/http"

	"Society-Synergy/base/services"

	"Society-Synergy/base/models"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) Login(ctx *gin.Context) {
	var user models.Login
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := uc.UserService.LoginUser(&user)
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

func (uc *UserController) Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("must provide all fields, email is missing")})
		return
	}

	user.Role = "STUDENT"
	valid := services.IsPasswordValid(user.Password)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password", "message": "Password must contain UPPER CASE, LOWER CASE, SPECIAL CHARACTER, NUMBER and LENGTH>7"})
		return
	}

	token, log, err := uc.UserService.RegisterUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data := map[string]string{
		"token": token,
	}
	// ctx.JSON(http.StatusOK, gin.H{"message": "Registration success", "data": data})
	ctx.Set("data", data)
	ctx.Set("log", &log)
	ctx.Next()
}
