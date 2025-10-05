package route

import (
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/handler"
	"github.com/gin-gonic/gin"
)

func SetupHealthcheckRoute(api *gin.Engine, handler handler.HealthcheckHandler) {
	router := api.Group("/healthcheck")
	{
		router.GET("", handler.Healthcheck)
	}
}