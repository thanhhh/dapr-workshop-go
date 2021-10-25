package http

import (
	vr "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration"

	"github.com/labstack/echo/v4"
)

func MapRoutes(httpGroup *echo.Group, h vr.Handlers) {
	httpGroup.GET("/vehicleinfo/:licenseNumber", h.GetVehicleInfo())
}
