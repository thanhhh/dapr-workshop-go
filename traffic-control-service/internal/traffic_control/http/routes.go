package http

import (
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"

	"github.com/labstack/echo/v4"
)

func MapRoutes(commGroup *echo.Group, h tc.Handlers) {
	commGroup.POST("entrycam", h.VehicleEntry())
	commGroup.POST("exitcam", h.VehicleExit())
}
