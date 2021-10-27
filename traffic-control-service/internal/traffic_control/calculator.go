package trafficcontrol

import "time"

type SpeedingViolationCalculator interface {
	DetermineSpeedingViolationInKmh(entryTimestamp time.Time, exitTimestamp time.Time) int
	GetRoadId() string
}
