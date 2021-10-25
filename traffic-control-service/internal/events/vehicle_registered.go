package events

import "time"

type VehicleRegistered struct {
	Lane          string    `json:"lane" validate:"required"`
	LicenseNumber string    `json:"license_number" validate:"required"`
	Timestamp     time.Time `json:"timestamp" validate:"required"`
}
