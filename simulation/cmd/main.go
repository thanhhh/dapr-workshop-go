package main

import (
	"log"

	"dapr-workshop-go/simulation/internal/proxies"
	"dapr-workshop-go/simulation/internal/simulation"
)

func main() {
	log.Println("Starting Simulator")

	for {
		go startCameraSimulationLane(1)
		go startCameraSimulationLane(2)
		go startCameraSimulationLane(3)
		go startCameraSimulationLane(4)
		go startCameraSimulationLane(5)
		go startCameraSimulationLane(6)
		select {}
	}
}

func startCameraSimulationLane(camNumber int) {
	service := proxies.NewService()
	cameraSimulator := simulation.NewSimulator(service, camNumber)
	cameraSimulator.Start()
}
