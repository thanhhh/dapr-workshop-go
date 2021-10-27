package trafficcontrol

import (
	"github.com/labstack/echo/v4"
)

type Handlers interface {
	VehicleEntry() echo.HandlerFunc
	VehicleExit() echo.HandlerFunc
}
