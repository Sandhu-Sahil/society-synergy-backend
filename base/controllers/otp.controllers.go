package controllers

import (
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) SendOtp(ctx *gin.Context) {
	user_id, _, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = uc.UserService.SendOTP(user_id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (uc *UserController) VerifyOtp(ctx *gin.Context) {
	user_id, _, err := token.ExtractTokenID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var otp struct {
		OTP string `json:"otp"`
	}
	if err := ctx.ShouldBindJSON(&otp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = uc.UserService.VerifyOTP(user_id, otp.OTP)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully"})
}
