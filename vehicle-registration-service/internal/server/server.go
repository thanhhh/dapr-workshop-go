package server

import (
	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/server"

	"github.com/labstack/echo/v4"

	vrHttp "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration/http"
	vrRepos "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration/repositories"
)

type vehicleRegistrationServer struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
}

func NewServerHandler(echo *echo.Echo, cfg *config.Config, logger logger.Logger) server.ServerHandlers {
	return &vehicleRegistrationServer{echo: echo, cfg: cfg, logger: logger}
}

func (s vehicleRegistrationServer) MapHandlers(e *echo.Echo) error {
	repository := vrRepos.NewInMemoryRepository()
	vrHandlers := vrHttp.NewHandlers(s.cfg, repository, s.logger)
	httpGroup := e.Group("/")
	vrHttp.MapRoutes(httpGroup, vrHandlers)
	return nil
}
