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

type defaultVehicleInfoService struct {
	logger logger.Logger
}

func NewProxy(logger logger.Logger) fc.VehicleInfoService {
	return &defaultVehicleInfoService{logger: logger}
}

func (p *defaultVehicleInfoService) GetVehicleInfo(licenseNumber string) (models.VehicleInfo, error) {
	vehicleInfo := models.VehicleInfo{}

	url := fmt.Sprintf("http://127.0.0.1:6002/vehicleinfo/%s", licenseNumber)

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
