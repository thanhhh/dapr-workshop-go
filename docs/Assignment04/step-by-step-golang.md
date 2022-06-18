# Assignment 4 - Add Dapr state management

## Assignment goals

To complete this assignment, you must reach the following goals:

- The TrafficControl service saves the state of a vehicle (`VehicleState` class) using the state management building
  block after vehicle entry.
- The TrafficControl service reads and updates the state of a vehicle using the state management
  building block after vehicle exit.

This assignment targets number **3** in the end-state setup:

<img src="../img/dapr-setup.png" style="zoom: 67%;" />

## Step 1: Use the Dapr state management building block

First, you need to add something to the state management configuration file:

1. Open the file `dapr/components/statestore.yaml` in VS Code.

1. Add a `scopes` section to the configuration file that specifies that only the TrafficControlService should use the
   state management building block:

   ```yaml
   apiVersion: dapr.io/v1alpha1
   kind: Component
   metadata:
     name: statestore
   spec:
     type: state.redis
     version: v1
     metadata:
     - name: redisHost
       value: localhost:6379
     - name: redisPassword
       value: ""
     - name: actorStateStore
       value: "true"
   scopes:
     - trafficcontrolservice
   ```

Now you will add code to the TrafficControlService so it uses the Dapr state management building block to store vehicle
state:

1. Open the file `traffic-control-service/internal/traffic_control/http/handlers.go` in VS Code.

2. Inspect the code in the `VehicleEntry` and `VehicleExit` methods of this handler. The methods refer to a `repository`
   which is defined earlier in the file:

   ```go
   repository := tcRepositories.NewVehicleStateRepository()
   ```
   You can see in package `traffic-control-service/internal/traffic_control/repositories`, we deined a memory repository.
   In this assignment, we will define new repository using Dapr state store instead.

3. Create the file `traffic-control-service/internal/traffic_control/repositories/state_repository.go` in VS Code 

4. Define state store name

   ```go
   const STATE_STORE_NAME = "statestore"
   ```

5. Define state repository struct 

   ```go
   type stateStoreRepository struct {
	   logger logger.Logger
   }
   ```

6. Define the entry state struct

   ```go
   type StateEntry struct {
	   Key   string              `json:"key"`
	   Value models.VehicleState `json:"value"`
   }
   ```

7. Provide factory function to create `stateStoreRepository`

   ```go
   func NewStateStoreRepository(logger logger.Logger) tc.VehicleStateRepository {
	   return &stateStoreRepository{logger: logger}
   }
   ```

8. Finally, implement for the interface repository 
`traffic-control-service/internal/traffic_control/repository.go`, `Save` method

    ```go
    func (r *stateStoreRepository) Save(state models.VehicleState) error {

      var entries [1]StateEntry

      entries[0] = StateEntry{
         Key:   state.LicenseNumber,
         Value: state,
      }

      entriesJSON, err := json.Marshal(entries)
      if err != nil {
        r.logger.Error(err)
        return fmt.Errorf("StateEntry encode json error: %v", err)
      }

      url := fmt.Sprintf("http://localhost:3600/v1.0/state/%s", STATE_STORE_NAME)

      resp, err := http.Post(url, "application/json", bytes.NewBuffer(entriesJSON))
      if err != nil {
        r.logger.Error(err)
        return fmt.Errorf("StateEntry send http to Dapr state store error: %v", err)
      }

      defer resp.Body.Close()

      return nil
    }    
    ```

   In the above, we create an entry state with license number as key and send it to Dapr state management API with HTTP post request.

9. Implement repository `Get` method to get state from Dapr state management API with HTTP get request

    ```go
    func (r stateStoreRepository) Get(licenseNumber string) (models.VehicleState, error) {
      vehicleState := models.VehicleState{}
      url := fmt.Sprintf("http://localhost:3600/v1.0/state/%s/%s", STATE_STORE_NAME, licenseNumber)

      resp, err := http.Get(url)
      if err != nil {
        r.logger.Error(err)

        return vehicleState, fmt.Errorf("Get http Dapr state store error: %v", err)
      }

      data, err := ioutil.ReadAll(resp.Body)
      if err != nil {
        r.logger.Fatal(err)

        return vehicleState, fmt.Errorf("read state store data error: %v", err)
      }

      err = json.Unmarshal(data, &vehicleState)
      if err != nil {
        r.logger.Fatal(err)

        return vehicleState, fmt.Errorf("parse json state store data error: %v", err)
      }
      return vehicleState, nil
    }
    ```
10. Open the file `traffic-control-service/internal/server/server.go`, go to `MapHandlers` method

11. Replace `inmemoryRepository` with new `stateStoreRepository`

    ```go
    repository := tcRepositories.NewStateStoreRepository(s.logger)
    ```

Now you're ready to test the application.

## Step 2a: Test the application

1. Make sure no services from previous tests are running (close the terminal windows)
2. Make sure all the Docker containers introduced in the previous assignments are running (you can use the
   `infrastructure/start-all.sh` script to start them).
3. Open the terminal window in VS Code and make sure the current folder is `VehicleRegistrationService`.
4. Enter the following command to run the VehicleRegistrationService with a Dapr sidecar:

   ```console
   dapr run --app-id vehicleregistrationservice \
   			  --app-port 6002 \
   			  --dapr-http-port 3602 \
   			  --dapr-grpc-port 60002 \
   			  --components-path ../dapr/components \
   			  go run ./cmd/main.go
   ```

5. Open a **new** terminal window in VS Code and change the current folder to `FineCollectionService`.

6. Enter the following command to run the FineCollectionService with a Dapr sidecar:

   ```console
   dapr run --app-id finecollectionservice \
   			  --app-port 6001 \
   			  --dapr-http-port 3601 \
   			  --dapr-grpc-port 60001 \
   			  --components-path ../dapr/components \
   			  go run ./cmd/main.go
   ```

7. Open a **new** terminal window in VS Code and change the current folder to `TrafficControlService`.

8. Enter the following command to run the TrafficControlService with a Dapr sidecar:

   ```console
   dapr run --app-id trafficcontrolservice \
   			  --app-port 6000 \
   			  --dapr-http-port 3600 \
   			  --dapr-grpc-port 60000 \
   			  --components-path ../dapr/components \
   			  go run ./cmd/main.go
   ```

9. Open a **new** terminal window in VS Code and change the current folder to `Simulation`.

10. Start the simulation:

    ```console
    go run ./cmd/main.go
    ```

You should see similar logging as before.

## Step 2b: Verify the state-store

Obviously, the behavior of the application is exactly the same as before. But are the VehicleState entries actually
stored in the default Redis state-store? To check this, you will use the redis CLI inside the `dapr_redis` container
that is used as state-store in the default Dapr installation.

1. Open a **new** terminal window in VS Code.

2. Execute the following command to start the redis-cli inside the running `dapr_redis` container:

   ```console
   docker exec -it dapr_redis redis-cli
   ```

3. In the redis-cli enter the following command to get the list of keys of items stored in the redis cache:

   ```console
   keys *
   ```

   You should see a list of entries with keys in the form `"trafficcontrolservice||<license-number>"`.

4. Enter the following command in the redis-cli to get the data stored with this key (change the license-number to one
   in the list you see):

   ```console
   hgetall trafficcontrolservice||KL-495-J
   ```

5. You should see something similar to this:

   ```console
   ❯ docker exec -it dapr_redis redis-cli
   127.0.0.1:6379> keys *
   1) "trafficcontrolservice||18-RSS-4"
   2) "trafficcontrolservice||84-GJ-06"
   3) "trafficcontrolservice||KJ-HS-06"
   4) "trafficcontrolservice||JN-TH-23"
   5) "trafficcontrolservice||11-GT-84"
   6) "trafficcontrolservice||RN-KR-35"
   7) "trafficcontrolservice||10-HYD-5"
   8) "trafficcontrolservice||YT-66-PY"
   9) "trafficcontrolservice||ND-841-Y"
   10) "trafficcontrolservice||T-375-NF"
   127.0.0.1:6379> hgetall trafficcontrolservice||ND-841-Y
   1) "version"
   2) "1"
   3) "data"
   4) "{\"licenseNumber\":\"ND-841-Y\",\"entryTimestamp\":\"2021-09-15T11:19:18.1781609+02:00\",\"exitTimestamp\":\"0001-01-01T00:00:00\"}"
   ```

As you can see, the data is actually stored in the redis cache. The cool thing about Dapr is that the state management
building block supports different state-stores through its component model. So without changing any code but only
specifying a different Dapr component configuration, you could use an entirely different storage mechanism.

> If you're up for it, try to swap-out Redis with another state provider. See the 
> [the list of available stores in the Dapr documentation](https://docs.dapr.io/operations/components/setup-state-store/supported-state-stores/)).
> To configure a different state-store, you need to change the file `dapr/components/statestore.yaml`.

## Step 3: Use Dapr state management with the Dapr SDK for Go

In this step you're going to change the `VehicleStateRepository` class and replace calling the Dapr state management
API directly over HTTP with using the `DaprClient` from the Dapr SDK for Python.

1. Open the file `TrafficControlService/traffic_control/repositories.py` in VS Code.

2. Add an import statement to the top of the file for the Dapr client:

   ```python
   from dapr.clients import DaprClient
   ```

3. Replace the the code in the `set_vehicle_state` with the following code:

   ```python
   with DaprClient() as client:
        client.save_state("statestore", vehicle_state.license_number, vehicle_state.json())
   ```

4. Replace the implementation of the `get_vehicle_state` method with the following code:

   ```python
   with DaprClient() as client:
        return models.VehicleState.parse_raw(client.get_state("statestore", license_number).text())
   ```

The repository code should now look like this:

```python
import requests
from . import models
from dapr.clients import DaprClient


class VehicleStateRepository:
    def __init__(self):
        self.state = {}

    def get_vehicle_state(self, license_number: str) -> models.VehicleState or None:
        with DaprClient() as client:
            return models.VehicleState.parse_raw(client.get_state("statestore", license_number).text())

    def set_vehicle_state(self, vehicle_state: models.VehicleState) -> None:
        with DaprClient() as client:
            client.save_state("statestore", vehicle_state.license_number, vehicle_state.json())
```

Now you're ready to test the application. Just repeat steps 2a and 2b.

## Next assignment

Make sure you stop all running processes and close all the terminal windows in VS Code before proceeding to the next assignment.

Go to [assignment 5](../Assignment05/README.md).
