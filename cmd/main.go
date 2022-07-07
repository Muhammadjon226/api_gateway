package main

import (
	"github.com/Muhammadjon226/api_gateway/api"
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/Muhammadjon226/api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "api_gateway")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Config{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
