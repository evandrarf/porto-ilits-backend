package main

import (
	"fmt"

	"github.com/evandrarf/porto-ilits-backend/internal/config"
	"github.com/evandrarf/porto-ilits-backend/internal/pkg/validate"
)

func main() {
  viperConfig:= config.NewViper()
	log := config.NewLogger(viperConfig)
	router:= config.NewApi(viperConfig, log)
	validator := validate.NewValidator()

	config.Bootstrap(&config.BootstrapConfig{
		Api:     router,
		Config:  viperConfig,
		Log:     log,
		DB:      nil, // Replace with actual DB instance if needed
		JWT:     nil, // Replace with actual JWT instance if needed
		Validator: validator,
	})

	listeningHost := viperConfig.GetString("api.host")
	listeningPort := viperConfig.GetString("api.port")

  router.Run(
		fmt.Sprintf("%s:%s", listeningHost, listeningPort),
	) 
}