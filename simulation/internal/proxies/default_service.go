package proxies

import (
	"encoding/json"
	"fmt"
	"log"

	tc "dapr-workshop-go/simulation/internal"
	"dapr-workshop-go/simulation/internal/events"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type defaultService struct {
	messagingAdapter mqtt.Client
}

func NewService(messagingAdapter mqtt.Client) tc.Service {
	return &defaultService{messagingAdapter: messagingAdapter}
}

func (s *defaultService) SendVehicleEntry(vehicleRegistered events.VehicleRegistered) error {
	var err error

	data, err := json.Marshal(vehicleRegistered)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleEntry encode json error: %v", err)
	}

	token := s.messagingAdapter.Publish("trafficcontrol/entrycam", 0, false, data)
	token.Wait()

	return nil
}

func (s *defaultService) SendVehicleExit(vehicleRegistered events.VehicleRegistered) error {
	var err error

	data, err := json.Marshal(vehicleRegistered)
	if err != nil {
		log.Print(err)

		return fmt.Errorf("SendVehicleExit encode json error: %v", err)
	}

	token := s.messagingAdapter.Publish("trafficcontrol/exitcam", 0, false, data)
	token.Wait()

	return nil
}
