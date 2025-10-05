package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	HealthcheckHandler interface {
		Healthcheck(ctx *gin.Context) 
	}

	healthcheckHandler struct{}
)

func NewHealthcheckHandler() HealthcheckHandler {
	return &healthcheckHandler{}
}

func (h *healthcheckHandler) Healthcheck(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"message": "Service is healthy",
	})
}