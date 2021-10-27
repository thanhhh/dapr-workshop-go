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

func NewProxy(logger logger.Logger) fc.VehicleInfoService {
	return &proxyVehicleInfoService{logger: logger}
}

func (p *proxyVehicleInfoService) GetVehicleInfo(vehicleId string) (models.VehicleInfo, error) {
	vehicleInfo := models.VehicleInfo{}

	url := fmt.Sprintf("http://localhost:3601/v1.0/invoke/vehicleregistrationservice/method/vehicleinfo/%s", vehicleId)

	resp, err := http.Get(url)
	if err != nil {
		p.logger.Error(err)

		return vehicleInfo, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		p.logger.Error(err)
		return vehicleInfo, err
	}
	if err := json.Unmarshal(body, &vehicleInfo); err != nil {
		p.logger.Error(err)
		return vehicleInfo, err
	}

	return vehicleInfo, nil
}
