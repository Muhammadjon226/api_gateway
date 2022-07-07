package v1

import (
	"github.com/Muhammadjon226/api_gateway/config"
	"github.com/Muhammadjon226/api_gateway/pkg/logger"
	"github.com/Muhammadjon226/api_gateway/services"
)

//HandlerV1 ...
type HandlerV1 struct {
	log            logger.Logger
	serviceManager services.IServiceManager
	cfg            config.Config
}


// New ...
func New(logger logger.Logger, serviceManager services.IServiceManager, cfg config.Config) *HandlerV1 {
	return &HandlerV1{
		log:            logger,
		serviceManager: serviceManager,
		cfg:            cfg,
	}
}
