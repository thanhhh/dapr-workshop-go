package trafficcontrol

import "dapr-workshop-go/simulation/internal/events"

type Service interface {
	SendVehicleEntry(vehicleRegistered events.VehicleRegistered) error
	SendVehicleExit(vehicleRegistered events.VehicleRegistered) error
}
