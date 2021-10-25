package models

import (
	"time"
)

type SpeedingViolation struct {
	LicenseNumber  string    `json:"license_number"`
	RoadId         string    `json:"road_id"`
	ViolationInKmh int       `json:"violation_in_kmh"`
	Timestamp      time.Time `json:"timestamp"`
}
