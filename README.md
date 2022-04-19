# Dapr workshop Golang

The Dapr workshop is a workshop that teaches you how to apply [Dapr](https://dapr.io) to a microservices application. This repository contains the source-code that forms the starting point of the Golang version of the workshop. With each workshop assignment you will expand the application with a Dapr building block.  

See the [Dapr workshop repository](https://github.com/edwinvw/dapr-workshop) for more information and instructions on how to get started.


### Running

1. Simulation

```
cd simulation
go run ./cmd/main.go
```

2. Vehicle Registration Service Dapr

```
cd vehicle-registration-service
dapr run --app-id vehicleregistrationservice --app-port 6003 --dapr-http-port 3602 --dapr-grpc-port 60003 go run ./cmd/main.go
```

3. Fine Collection Service Dapr

```
cd fine-collection-service
dapr run --app-id finecollectionservice --app-port 6001 --dapr-http-port 3601 --dapr-grpc-port 60001 go run ./cmd/main.go
```

4. Traffic Control Service

```
cd fine-collection-servie
go run ./cmd/main.go
```