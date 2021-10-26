package models

import (
	"time"
)

type SpeedingViolation struct {
	VehicleId      string    `json:"vehicle_id"`
	RoadId         string    `json:"road_id"`
	ViolationInKmh int       `json:"violation_in_kmh"`
	Timestamp      time.Time `json:"timestamp"`
}
