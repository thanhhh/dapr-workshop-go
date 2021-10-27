package vehicleregistration

import "dapr-workshop-go/vehicle-registration-service/internal/models"

type VehicleInfoRepository interface {
	Get(licenseNumber string) models.VehicleInfo
}
