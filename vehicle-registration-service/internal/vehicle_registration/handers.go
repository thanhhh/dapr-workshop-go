package vehicleregistration

import "github.com/labstack/echo/v4"

type Handlers interface {
	GetVehicleInfo() echo.HandlerFunc
}
