package vehicle_registration

import "dapr-workshop-go/vehicle-registration-service/internal/models"

type VehicleInfoRepository interface {
	Get(licenseNumber string) models.VehicleInfo
}
