package services

import (
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"
	"math"
	"time"
)

type defaultSpeedingViolationCalculator struct {
	roadId               string
	sectionLengthInKm    int
	maxAllowedSpeedInKmh int
	legalCorrectInKmh    int
}

func NewSpeedingViolationCalculator(
	roadId string,
	sectionLengthInKm int,
	maxAllowedSpeedInKmh int,
	legalCorrectInKmh int) tc.SpeedingViolationCalculator {
	return &defaultSpeedingViolationCalculator{
		roadId:               roadId,
		sectionLengthInKm:    sectionLengthInKm,
		maxAllowedSpeedInKmh: maxAllowedSpeedInKmh,
		legalCorrectInKmh:    legalCorrectInKmh,
	}
}

func (d defaultSpeedingViolationCalculator) DetermineSpeedingViolationInKmh(entryTimestamp time.Time, exitTimestamp time.Time) int {
	elapsedMinutes := exitTimestamp.Sub(entryTimestamp).Seconds()
	avgSpeedInKmh := math.Round((float64(d.sectionLengthInKm) / elapsedMinutes) * 60)
	violation := int(avgSpeedInKmh - float64(d.maxAllowedSpeedInKmh) - float64(d.legalCorrectInKmh))
	return violation
}

func (d defaultSpeedingViolationCalculator) GetRoadId() string {
	return d.roadId
}
