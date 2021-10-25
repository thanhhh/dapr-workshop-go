package server

import (
	"github.com/labstack/echo/v4"

	tcHttp "dapr-workshop-go/traffic-control-service/internal/traffic_control/http"
	tcRepositories "dapr-workshop-go/traffic-control-service/internal/traffic_control/repositories"
	tcServices "dapr-workshop-go/traffic-control-service/internal/traffic_control/services"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	svcService := tcServices.NewSpeedingViolationCalculator("A12", 10, 100, 5)
	vsRepository := tcRepositories.NewVehicleStateRepository()
	tcHandlers := tcHttp.NewTrafficControlHandlers(s.cfg, svcService, vsRepository, s.logger)
	tcApiGroup := e.Group("/")
	tcHttp.MapTrafficControlRoutes(tcApiGroup, tcHandlers)
	return nil
}
