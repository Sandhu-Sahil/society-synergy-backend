package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (c *UserController) Ping(ctx *gin.Context) {
	// sleep for 1 second
	// time.Sleep(2 * time.Second)
	ctx.JSON(http.StatusOK, gin.H{"message": "ping return pong"})
}
