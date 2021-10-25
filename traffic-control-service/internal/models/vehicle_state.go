package models

import (
	"time"
)

type VehicleState struct {
	LicenseNumber  string    `json:"license_number" validate:"required"`
	EntryTimestamp time.Time `json:"entry_timestamp" validate:"required"`
	ExitTimestamp  time.Time `json:"exit_timestamp"`
}
