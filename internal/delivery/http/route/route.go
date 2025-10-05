package route

import (
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/handler"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
)

type RouteConfig struct {
	Api                    *gin.Engine
	Middleware             *middleware.Middleware
	HealthcheckHandler    handler.HealthcheckHandler
	PostHandler           handler.PostHandler
}

func Setup(c *RouteConfig) {
	c.Api.Use(gin.Recovery())
	c.Api.Use(c.Middleware.CorsMiddleware())

	SetupHealthcheckRoute(c.Api, c.HealthcheckHandler)
	SetupPostRoute(c.Api, c.PostHandler, c.Middleware)
}