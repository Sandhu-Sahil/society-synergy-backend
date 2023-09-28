package controllers

import (
	"Society-Synergy/base/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (lc *LogsController) RegisterLog(ctx *gin.Context) {
	log := ctx.MustGet("log").(*models.AuditLogs)
	_, err := lc.LogsService.RegisterLog(log)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Registration success", "data": ctx.MustGet("data").(map[string]string)})
}
