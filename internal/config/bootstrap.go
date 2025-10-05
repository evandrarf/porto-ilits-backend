package config

import (
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/handler"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/middleware"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/repository"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/route"
	"github.com/evandrarf/porto-ilits-backend/internal/delivery/http/usecase"
	"github.com/evandrarf/porto-ilits-backend/internal/pkg/validate"
	"github.com/evandrarf/porto-ilits-backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type BootstrapConfig struct {
	Api       *gin.Engine
	Config    *viper.Viper
	DB        *gorm.DB
	Log       *logrus.Logger
	Validator *validate.Validator
	JWT       *jwt.JWT
}

func Bootstrap(config *BootstrapConfig) {
	mid:= middleware.NewMiddleware(&middleware.MiddlewareConfig{
		Log:    config.Log,
		JWT:    config.JWT,
		Config: config.Config,
		// AuthUsecase: nil, // Replace with actual usecase if needed
	})

	postRepository := repository.NewPostRepository(config.DB)

	postUsecase := usecase.NewPostUsecase(config.DB, postRepository)

	healthcheckHandler := handler.NewHealthcheckHandler()
	postHandler := handler.NewPostHandler(config.Validator, config.Log, postUsecase)

	route.Setup(&route.RouteConfig{
		Api:        config.Api,
		Middleware: mid,
		HealthcheckHandler: healthcheckHandler,
		PostHandler: postHandler,
	})
}