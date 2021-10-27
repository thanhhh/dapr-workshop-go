package utils

import (
	"github.com/labstack/echo/v4"

	"dapr-workshop-go/pkg/logger"
)

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config_docker.yml"
	}
	return "./config/config_local.yml"
}

func GetRequestID(c echo.Context) string {
	return c.Response().Header().Get(echo.HeaderXRequestID)
}
func GetIPAddress(c echo.Context) string {
	return c.Request().RemoteAddr
}

func LogResponseError(ctx echo.Context, logger logger.Logger, err error) {
	logger.Errorf(
		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}
