package repositories

import (
	"context"

	"encoding/json"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"

	"dapr-workshop-go/pkg/errors"
	"dapr-workshop-go/pkg/logger"
	models "dapr-workshop-go/traffic-control-service/internal/models"
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"
)

const STATE_STORE_NAME = "statestore"

type stateStoreRepository struct {
	logger logger.Logger
}

type StateEntry struct {
	Key   string              `json:"key"`
	Value models.VehicleState `json:"value"`
}

func NewStateStoreRepository(logger logger.Logger) tc.VehicleStateRepository {
	return &stateStoreRepository{logger: logger}
}

func (r *stateStoreRepository) Save(state models.VehicleState) error {
	var entries [1]StateEntry

	entries[0] = StateEntry{
		Key:   state.LicenseNumber,
		Value: state,
	}

	entryJSON, err := json.Marshal(entries)
	if err != nil {
		r.logger.Error(err)
		return fmt.Errorf("StateEntry encode json error: %v", err)
	}

	client, err := dapr.NewClient()

	if err != nil {
		r.logger.Error(err)
		return fmt.Errorf("create Dapr client error: %v", err)
	}

	ctx := context.Background()

	err = client.SaveState(ctx, STATE_STORE_NAME, state.LicenseNumber, []byte(entryJSON), nil)
	if err != nil {
		r.logger.Error(err)
		return fmt.Errorf("save state using Dapr client error: %v", err)
	}

	return nil
}

func (r stateStoreRepository) Get(licenseNumber string) (models.VehicleState, error) {
	client, err := dapr.NewClient()

	if err != nil {
		r.logger.Error(err)
		return models.VehicleState{}, fmt.Errorf("create Dapr client error: %v", err)
	}

	ctx := context.Background()
	result, err := client.GetState(ctx, STATE_STORE_NAME, licenseNumber, nil)

	if err != nil {
		r.logger.Error(err)
		return models.VehicleState{}, fmt.Errorf("get state from Daprt client error: %v", err)
	}

	if len(result.Value) == 0 {
		return models.VehicleState{}, errors.ErrVehicleStateRecordNotFound
	}

	var entries []StateEntry

	err = json.Unmarshal(result.Value, &entries)
	if err != nil {
		r.logger.Fatal(err)

		return models.VehicleState{}, fmt.Errorf("parse json state store data error: %v", err)
	}

	if len(entries) == 0 {
		return models.VehicleState{}, errors.ErrVehicleStateRecordNotFound
	}
	return entries[0].Value, nil
}
