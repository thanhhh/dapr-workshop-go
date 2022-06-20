package simulation

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"time"

	tc "dapr-workshop-go/simulation/internal"
	"dapr-workshop-go/simulation/internal/events"
)

type cameraSimulation struct {
	service   tc.Service
	camNumber int
}

const (
	minEntryDelayInMS = 50
	maxEntryDelayInMS = 5000
	minExitDelayInS   = 4
	maxExitDelayInS   = 10
)

var (
	validLicenseNumberChars = []rune("DFGHJKLNPRSTXYZ")
)

func NewSimulator(service tc.Service, camNumber int) tc.Simulation {
	return &cameraSimulation{service: service, camNumber: camNumber}
}

func (s *cameraSimulation) Start() {
	log.Printf("Start camera %d simulation.\n", s.camNumber)

	for {
		<-time.After(time.Duration(randInt(minEntryDelayInMS, maxEntryDelayInMS)) * time.Millisecond)
		s.generateEntry()
	}
}

func (s cameraSimulation) generateEntry() {
	entryTimestamp := time.Now()
	vehicleRegistered := events.VehicleRegistered{
		Lane:          s.camNumber,
		LicenseNumber: generateLicenseNumber(),
		Timestamp:     entryTimestamp,
	}

	s.service.SendVehicleEntry(vehicleRegistered)

	log.Printf(
		"Simulated ENTRY of vehicle with license-number %s in lane %d.\n",
		vehicleRegistered.LicenseNumber,
		vehicleRegistered.Lane)

	// simulate exit
	exitDelay := time.Duration(randInt(minExitDelayInS, maxExitDelayInS))
	time.Sleep(exitDelay * time.Second)

	vehicleRegistered.Timestamp = time.Now()
	vehicleRegistered.Lane = randInt(1, 4)
	s.service.SendVehicleExit(vehicleRegistered)

	log.Printf(
		"Simulated EXIT of vehicle with license-number %s in lane %d.\n",
		vehicleRegistered.LicenseNumber,
		vehicleRegistered.Lane)
}

func randInt(min, max int) int {
	return rand.Intn(max-min) + min
}

func generateLicenseNumber() string {
	typeLN := randInt(1, 9)
	kenteken := ""
	switch typeLN {
	case 1: // 99-AA-99
		kenteken = fmt.Sprintf("%02d-%s-%02d", randInt(1, 99), generateRandomCharacters(2), randInt(1, 99))
	case 2: // AA-99-AA
		kenteken = fmt.Sprintf("%s-%02d-%s", generateRandomCharacters(2), randInt(1, 99), generateRandomCharacters(2))
	case 3: // AA-AA-99
		kenteken = fmt.Sprintf("%s-%s-%02d", generateRandomCharacters(2), generateRandomCharacters(2), randInt(1, 99))
	case 4: // 99-AA-AA
		kenteken = fmt.Sprintf("%02d-%s-%s", randInt(1, 99), generateRandomCharacters(2), generateRandomCharacters(2))
	case 5: // 99-AAA-9
		kenteken = fmt.Sprintf("%02d-%s-%d", randInt(1, 99), generateRandomCharacters(3), randInt(1, 9))
	case 6: // 9-AAA-99
		kenteken = fmt.Sprintf("%d-%s-%02d", randInt(1, 9), generateRandomCharacters(3), randInt(1, 99))
	case 7: // AA-999-A
		kenteken = fmt.Sprintf("%s-%03d-%s", generateRandomCharacters(2), randInt(1, 999), generateRandomCharacters(1))
	case 8: // A-999-AA
		kenteken = fmt.Sprintf("%s-%03d-%s", generateRandomCharacters(1), randInt(1, 999), generateRandomCharacters(2))
	}

	return kenteken
}

func generateRandomCharacters(aantal int) string {
	var buffer bytes.Buffer

	for i := 0; i < aantal; i++ {
		randomIndex := rand.Intn(len(validLicenseNumberChars))
		buffer.WriteRune(validLicenseNumberChars[randomIndex])
	}
	return buffer.String()
}
