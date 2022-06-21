package proxies

import (
	"context"
	"encoding/json"
	"fmt"

	dapr "github.com/dapr/go-sdk/client"

	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
	"dapr-workshop-go/pkg/logger"
)

type defaultVehicleInfoService struct {
	logger logger.Logger
}

func NewProxy(logger logger.Logger) fc.VehicleInfoService {
	return &defaultVehicleInfoService{logger: logger}
}

func (p *defaultVehicleInfoService) GetVehicleInfo(
	ctx context.Context, vehicleId string) (models.VehicleInfo, error) {
	vehicleInfo := models.VehicleInfo{}

	daprClient, err := dapr.NewClient()
	if err != nil {
		p.logger.Error(err)
		return vehicleInfo, fmt.Errorf("create dapr client error %v", err)
	}

	methodName := fmt.Sprintf("vehicleinfo/%s", vehicleId)
	resp, err := daprClient.InvokeMethod(ctx, "vehicleregistrationservice", methodName, "get")

	if err != nil {
		p.logger.Error(err)
		return vehicleInfo, fmt.Errorf("invoke method in dapr client error %v", err)
	}

	if err := json.Unmarshal(resp, &vehicleInfo); err != nil {
		p.logger.Error(err)
		return vehicleInfo, fmt.Errorf("decode json vehicle info error %v", err)
	}

	return vehicleInfo, nil
}
