package http

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/errors"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/utils"

	vr "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration"
)

type vehicleRegistrationHandlers struct {
	cfg        *config.Config
	repository vr.VehicleInfoRepository
	logger     logger.Logger
}

func NewHandlers(cfg *config.Config, repository vr.VehicleInfoRepository, logger logger.Logger) vr.Handlers {
	return &vehicleRegistrationHandlers{
		cfg:        cfg,
		repository: repository,
		logger:     logger,
	}
}

func (h *vehicleRegistrationHandlers) GetVehicleInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		licenseNumber := c.Param("licenseNumber")

		if licenseNumber == "" {
			utils.LogResponseError(c, h.logger, errors.NewBadRequestError("licenseNumber is required"))
			return c.JSON(http.StatusBadRequest, errors.NewBadRequestError("licenseNumber is required"))
		}

		vehicleInfo := h.repository.Get(licenseNumber)

		h.logger.Infof(
			"Requested vehicle registration information for license-number %s.", licenseNumber)
		return c.JSON(http.StatusOK, vehicleInfo)
	}
}
