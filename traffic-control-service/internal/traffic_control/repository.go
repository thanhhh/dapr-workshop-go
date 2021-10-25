package traffic_control

import "dapr-workshop-go/traffic-control-service/internal/models"

type VehicleStateRepository interface {
	Save(state models.VehicleState) error
	Get(licenseNumber string) (models.VehicleState, error)
}
