package finecollection

import "dapr-workshop-go/fine-collection-service/internal/models"

type VehicleInfoService interface {
	GetVehicleInfo(licenseNumber string) (models.VehicleInfo, error)
}
