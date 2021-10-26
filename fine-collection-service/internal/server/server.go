package server

import (
	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/server"

	"github.com/labstack/echo/v4"

	fcHttp "dapr-workshop-go/fine-collection-service/internal/fine_collection/http"
	fcProxies "dapr-workshop-go/fine-collection-service/internal/fine_collection/proxies"
	fcServices "dapr-workshop-go/fine-collection-service/internal/fine_collection/services"
)

type fineCollectionServer struct {
	echo   *echo.Echo
	cfg    *config.Config
	logger logger.Logger
}

func NewServerHandler(echo *echo.Echo, cfg *config.Config, logger logger.Logger) server.ServerHandlers {
	return &fineCollectionServer{echo: echo, cfg: cfg, logger: logger}
}

func (s *fineCollectionServer) MapHandlers(e *echo.Echo) error {
	vehicleService := fcProxies.NewProxy()
	emailService := fcServices.NewEmailService(s.logger)
	fineCalculator := fcServices.NewFineCalculator()

	fcHandlers := fcHttp.NewHandlers(s.cfg, s.logger, vehicleService, emailService, fineCalculator)

	httpGroup := e.Group("/")
	fcHttp.MapRoutes(httpGroup, fcHandlers)
	return nil
}
