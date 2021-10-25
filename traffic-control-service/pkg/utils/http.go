package utils

import (
	"context"
	"dapr-workshop-go/traffic-control-service/internal/events"
	"dapr-workshop-go/traffic-control-service/pkg/errors"
)

type VehicleRegisteredCtxKey struct{}

// Get user from context
func GetVehicleRegisteredFromCtx(ctx context.Context) (*events.VehicleRegistered, error) {
	value, ok := ctx.Value(VehicleRegisteredCtxKey{}).(*events.VehicleRegistered)

	if !ok {
		return nil, errors.BadRequest
	}

	return value, nil
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}
