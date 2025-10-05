package middleware

import (
	"github.com/evandrarf/porto-ilits-backend/pkg/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MiddlewareConfig struct {
	Log    *logrus.Logger
	JWT    *jwt.JWT
	Config *viper.Viper
	// usecase.AuthUsecase
}

type Middleware struct {
	Log    *logrus.Logger
	JWT    *jwt.JWT
	Config *viper.Viper
	// usecase.AuthUsecase
}

func NewMiddleware(c *MiddlewareConfig) *Middleware {
	return &Middleware{
		Log:         c.Log,
		JWT:         c.JWT,
		Config:      c.Config,
		// AuthUsecase: c.AuthUsecase,
	}
}
