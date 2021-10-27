package repositories

import (
	"dapr-workshop-go/pkg/errors"

	models "dapr-workshop-go/traffic-control-service/internal/models"
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"
)

type inmemoryRepository struct {
	states map[string]models.VehicleState
}

func NewVehicleStateRepository() tc.VehicleStateRepository {
	m := make(map[string]models.VehicleState)
	return &inmemoryRepository{states: m}
}

func (r *inmemoryRepository) Save(state models.VehicleState) error {
	r.states[state.LicenseNumber] = state
	return nil
}

func (r inmemoryRepository) Get(licenseNumber string) (models.VehicleState, error) {
	state, exists := r.states[licenseNumber]

	if exists {
		return state, nil
	}
	return models.VehicleState{}, errors.ErrVehicleStateRecordNotFound
}
