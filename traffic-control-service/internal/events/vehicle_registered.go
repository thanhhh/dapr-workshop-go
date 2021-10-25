package events

import "time"

type VehicleRegistered struct {
	Lane          int
	LicenseNumber string
	Timestamp     time.Time
}
