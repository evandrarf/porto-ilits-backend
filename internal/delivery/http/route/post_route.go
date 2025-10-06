package route

import (
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/handler"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

func SetupPostRoute(api *gin.Engine, handler handler.PostHandler, m *middleware.Middleware) {
	router := api.Group("/posts")
	{
		router.POST("", handler.Create)
		router.GET("", handler.GetAll)
		router.PUT("/:id", handler.Update)
		router.DELETE("/:id", handler.Delete)
	}
}