package finecollection

import (
	"context"
	"dapr-workshop-go/fine-collection-service/internal/models"
)

type VehicleInfoService interface {
	GetVehicleInfo(ctx context.Context, licenseNumber string) (models.VehicleInfo, error)
}
