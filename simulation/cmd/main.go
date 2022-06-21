package main

import (
	"log"
	"time"

	"dapr-workshop-go/simulation/internal/proxies"
	"dapr-workshop-go/simulation/internal/simulation"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	log.Println("Starting Simulator")

	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("simulation")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	for {
		go startCameraSimulationLane(c, 1)
		go startCameraSimulationLane(c, 2)
		go startCameraSimulationLane(c, 3)
		select {}
	}
}

func startCameraSimulationLane(mqttClient mqtt.Client, camNumber int) {
	service := proxies.NewService(mqttClient)
	cameraSimulator := simulation.NewSimulator(service, camNumber)
	cameraSimulator.Start()
}
