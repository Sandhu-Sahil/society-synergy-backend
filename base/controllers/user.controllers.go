package controllers

import (
	"Society-Synergy/base/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *Controller) GetUser(ctx *gin.Context) {
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
