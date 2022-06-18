package repositories

import (
	"bytes"
	"dapr-workshop-go/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

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

	url := fmt.Sprintf("http://localhost:3600/v1.0/state/%s", STATE_STORE_NAME)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(entryJSON))
	if err != nil {
		r.logger.Error(err)
		return fmt.Errorf("StateEntry send http to Dapr state store error: %v", err)
	}

	defer resp.Body.Close()

	return nil
}

func (r stateStoreRepository) Get(licenseNumber string) (models.VehicleState, error) {
	url := fmt.Sprintf("http://localhost:3600/v1.0/state/%s/%s", STATE_STORE_NAME, licenseNumber)

	resp, err := http.Get(url)
	if err != nil {
		r.logger.Error(err)

		return models.VehicleState{}, fmt.Errorf("Get http Dapr state store error: %v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		r.logger.Fatal(err)

		return models.VehicleState{}, fmt.Errorf("read state store data error: %v", err)
	}

	var vehicleState models.VehicleState

	err = json.Unmarshal(data, &vehicleState)
	if err != nil {
		r.logger.Fatal(err)

		return models.VehicleState{}, fmt.Errorf("parse json state store data error: %v", err)
	}

	// r.logger.Infof("%v", stateEntry)
	return vehicleState, nil
}
