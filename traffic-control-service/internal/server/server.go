package server

import (
	"github.com/labstack/echo/v4"

	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/server"

	tcHttp "dapr-workshop-go/traffic-control-service/internal/traffic_control/http"
	tcRepositories "dapr-workshop-go/traffic-control-service/internal/traffic_control/repositories"
	tcServices "dapr-workshop-go/traffic-control-service/internal/traffic_control/services"
)

type trafficControlServer struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
}

func NewServerHandler(echo *echo.Echo, cfg *config.Config, logger logger.Logger) server.ServerHandlers {
	return &trafficControlServer{echo: echo, cfg: cfg, logger: logger}
}

func (s *trafficControlServer) MapHandlers(e *echo.Echo) error {
	calculator := tcServices.NewSpeedingViolationCalculator("A12", 37, 100, 5)
	repository := tcRepositories.NewStateStoreRepository(s.logger)
	handlers := tcHttp.NewHandlers(s.cfg, calculator, repository, s.logger)
	apiGroup := e.Group("/")
	tcHttp.MapRoutes(apiGroup, handlers)
	return nil
}
