package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"

	"dapr-workshop-go/simulation/internal/proxies"
	"dapr-workshop-go/simulation/internal/simulation"
)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}

func main() {
	log.Println("Starting Simulator")

	opts := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883").SetClientID("simulation")
	opts.SetKeepAlive(2 * time.Second)
	opts.SetDefaultPublishHandler(f)
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

	time.Sleep(6 * time.Second)

	c.Disconnect(250)

	time.Sleep(1 * time.Second)
}

func startCameraSimulationLane(mqttClient mqtt.Client, camNumber int) {
	service := proxies.NewService(mqttClient)
	cameraSimulator := simulation.NewSimulator(service, camNumber)
	cameraSimulator.Start()
}
