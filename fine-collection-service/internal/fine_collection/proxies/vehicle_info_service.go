package proxies

import (
	"context"
	"encoding/json"

	dapr "github.com/dapr/go-sdk/client"

	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
	"dapr-workshop-go/pkg/logger"
)

type defaultVehicleInfoService struct {
	logger logger.Logger
}

func NewProxy(logger logger.Logger) fc.VehicleInfoService {
	return &defaultVehicleInfoService{
		logger: logger,
	}
}

func (p *defaultVehicleInfoService) GetVehicleInfo(ctx context.Context, licenseNumber string) (models.VehicleInfo, error) {
	vehicleInfo := models.VehicleInfo{}

	daprClient, err := dapr.NewClient()
	if err != nil {
		p.logger.Error(err)
	}
	methodName := "vehicleinfo/" + licenseNumber
	resp, err := daprClient.InvokeMethod(ctx, "vehicleregistrationservice", methodName, "get")

	if err != nil {
		p.logger.Error(err)
		return vehicleInfo, err
	}

	if err := json.Unmarshal(resp, &vehicleInfo); err != nil {
		p.logger.Error(err)
		return vehicleInfo, err
	}

	return vehicleInfo, nil
}
