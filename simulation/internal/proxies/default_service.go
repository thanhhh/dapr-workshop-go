package proxies

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	tc "dapr-workshop-go/simulation/internal"
	"dapr-workshop-go/simulation/internal/events"
)

type defaultService struct {
}

func NewService() tc.Service {
	return &defaultService{}
}

func (s *defaultService) SendVehicleEntry(vehicleRegistered events.VehicleRegistered) error {
	var err error

	data, err := json.Marshal(vehicleRegistered)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleEntry encode json error: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:6000/entrycam", bytes.NewBuffer(data))

	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleEntry create http request entrycam error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleEntry send http entrycam error: %v", err)
	}

	defer resp.Body.Close()

	return nil
}

func (s *defaultService) SendVehicleExit(vehicleRegistered events.VehicleRegistered) error {
	var err error

	data, err := json.Marshal(vehicleRegistered)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleExit encode json error: %v", err)
	}

	req, err := http.NewRequest("POST", "http://localhost:6000/exitcam", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleExit send http entrycam error: %v", err)
	}

	defer resp.Body.Close()

	return nil
}
