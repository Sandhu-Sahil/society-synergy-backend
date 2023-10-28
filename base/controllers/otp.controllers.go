package controllers

import (
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) EmailSendOtp(ctx *gin.Context) {
	var email struct {
		Email string `json:"email"`
	}
	if err := ctx.ShouldBindJSON(&email); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := uc.UserService.EmailSendOTP(email.Email)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}

func (uc *UserController) VerifyOtpEmail(ctx *gin.Context) {
	var otp struct {
		Email string `json:"email"`
		Otp   string `json:"otp"`
	}
	if err := ctx.ShouldBindJSON(&otp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := uc.UserService.VerifyOTPEmail(otp.Email, otp.Otp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data := map[string]interface{}{
		"token": token,
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully", "data": data})
}

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
