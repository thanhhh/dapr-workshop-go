package proxies

import (
	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
	"dapr-workshop-go/pkg/logger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type proxyVehicleInfoService struct {
	logger logger.Logger
}

func NewProxy() fc.VehicleInfoService {
	return &proxyVehicleInfoService{}
}

func (p proxyVehicleInfoService) GetVehicleInfo(vehicleId string) (models.VehicleInfo, error) {
	vehicleInfo := models.VehicleInfo{}

	url := fmt.Sprintf("http://127.0.0.1:6002/vehicleinfo/%s", vehicleId)

	resp, err := http.Get(url)
	if err != nil {
		p.logger.Fatal(err)

		return vehicleInfo, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.logger.Fatal(err)
		return vehicleInfo, err
	}
	if err := json.Unmarshal(body, &vehicleInfo); err != nil {
		p.logger.Fatal(err)
		return vehicleInfo, err
	}

	return vehicleInfo, nil
}
