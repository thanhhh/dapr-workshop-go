package http

import (
	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/utils"
	"fmt"
	"net/http"

	fc "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
	httpErrors "dapr-workshop-go/pkg/errors"

	"github.com/labstack/echo/v4"
)

type fineCollectionHandlers struct {
	cfg            *config.Config
	vehicleService fc.VehicleInfoService
	emailService   fc.EmailService
	calculator     fc.FineCalculator
	logger         logger.Logger
}

func NewHandlers(
	cfg *config.Config,
	logger logger.Logger,
	vehicleService fc.VehicleInfoService,
	emailService fc.EmailService,
	calculator fc.FineCalculator) fc.Handlers {
	return &fineCollectionHandlers{
		cfg:            cfg,
		vehicleService: vehicleService,
		emailService:   emailService,
		calculator:     calculator,
		logger:         logger,
	}
}

func (h *fineCollectionHandlers) CollectFine() echo.HandlerFunc {
	return func(c echo.Context) error {
		speedingViolation := &models.SpeedingViolation{}

		if err := c.Bind(speedingViolation); err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(http.StatusBadRequest, err)
		}

		if err := utils.ValidateStruct(c.Request().Context(), speedingViolation); err != nil {
			h.logger.Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}

		fine, err := h.calculator.CalculateFine(
			h.cfg.LicenseKey.FineCalculatorLicenseKey,
			speedingViolation.ViolationInKmh)

		if err := utils.ValidateStruct(c.Request().Context(), speedingViolation); err != nil {
			h.logger.Error(err)
			return c.JSON(http.StatusBadRequest, err)
		}

		// get owner info
		vehicleInfo, err := h.vehicleService.GetVehicleInfo(speedingViolation.VehicleId)

		if err != nil {
			h.logger.Error(err)
			return c.JSON(http.StatusBadRequest, httpErrors.NewBadRequestError(
				fmt.Sprintf("Vehicle Id %s is not found", speedingViolation.VehicleId)))
		}

		// log fine

		fineString := "tbd by the prosecutor"
		if fine > 0 {
			fineString = fmt.Sprintf("%d Euro", fine)
		}

		h.logger.Infof("Sent speeding ticket to %s. "+
			"Road: %s, Licensenumber: %s, "+
			"Vehicle: %s %s, "+
			"Violation: %d Km/h, Fine: %s, "+
			"On: %s at %s.",
			vehicleInfo.OwnerName,
			speedingViolation.RoadId,
			speedingViolation.VehicleId,
			vehicleInfo.Brand,
			vehicleInfo.Model,
			speedingViolation.ViolationInKmh,
			fineString,
			speedingViolation.Timestamp.Format("31-12-2021"),
			speedingViolation.Timestamp.Format("13:59:59"),
		)

		// send fine by email
		// TODO

		return c.NoContent(http.StatusOK)
	}
}