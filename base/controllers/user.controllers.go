package controllers

import (
	"Society-Synergy/base/models"
	"Society-Synergy/base/services"
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	user_id, _, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if id == "0" {
		id = user_id
	}

	u, err := uc.UserService.GetUserByID(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}

func (uc *UserController) ChangePassword(ctx *gin.Context) {
	var password struct {
		Otp         string `json:"otp"`
		NewPassword string `json:"password"`
	}
	if err := ctx.ShouldBindJSON(&password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	valid := services.IsPasswordValid(password.NewPassword)
	if !valid {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password", "message": "Password must contain UPPER CASE, LOWER CASE, SPECIAL CHARACTER, NUMBER and LENGTH>7"})
		return
	}

	user_id, _, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, err := uc.UserService.ChangePassword(user_id, password.Otp, password.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.RegisterLogS(&log)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user_id, _, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log, err := uc.UserService.UpdateUser(user_id, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = services.RegisterLogS(&log)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}
