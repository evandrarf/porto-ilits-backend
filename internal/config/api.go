package config

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func NewApi(config *viper.Viper, log *logrus.Logger) *gin.Engine {
	router := gin.New()

	router.Use(gin.Recovery())
	router.Use(RequestLogger(log))
	router.Use(ErrorHandler(log))

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}	

	return router
}

func RequestLogger(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := c.Request.URL.Path
		method := c.Request.Method
		c.Next()
		status := c.Writer.Status()

		log.WithFields(logrus.Fields{
			"status": status,
			"method": method,
			"path":   path,
		}).Info("HTTP request")
	}
}

func ErrorHandler(log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// Ambil error jika ada
		errs := c.Errors
		if len(errs) == 0 {
			return
		}

		err := errs.Last().Err
		statusCode := c.Writer.Status()
		if statusCode < 400 {
			statusCode = http.StatusInternalServerError
		}

		if statusCode >= 500 {
			log.WithError(err).Error("Internal server error")
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Internal Server Error",
			})
			return
		}

		log.WithError(err).Warn("Request failed")
		c.JSON(statusCode, gin.H{
			"success": false,
			"message": err.Error(),
		})
	}
}