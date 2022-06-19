package http

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	dapr "github.com/dapr/go-sdk/client"

	"github.com/labstack/echo/v4"

	"dapr-workshop-go/pkg/config"
	"dapr-workshop-go/pkg/logger"
	"dapr-workshop-go/pkg/utils"

	"dapr-workshop-go/traffic-control-service/internal/events"
	"dapr-workshop-go/traffic-control-service/internal/models"
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"
)

type trafficControlHandlers struct {
	cfg        *config.Config
	service    tc.SpeedingViolationCalculator
	repository tc.VehicleStateRepository
	logger     logger.Logger
}

func NewHandlers(cfg *config.Config, service tc.SpeedingViolationCalculator, repository tc.VehicleStateRepository, logger logger.Logger) tc.Handlers {
	return &trafficControlHandlers{cfg: cfg, service: service, repository: repository, logger: logger}
}

func (h *trafficControlHandlers) VehicleEntry() echo.HandlerFunc {
	return func(c echo.Context) error {
		message := &events.VehicleRegistered{}

		if err := c.Bind(message); err != nil {
			h.logger.Error(err)
			return c.NoContent(http.StatusBadRequest)
		}

		if err := utils.ValidateStruct(c.Request().Context(), message); err != nil {
			h.logger.Error(err)
			return c.NoContent(http.StatusBadRequest)
		}

		h.logger.Infof(
			"ENTRY detected in lane %d at %s of vehicle with license-number %s.",
			message.Lane, message.Timestamp.Format("15:04:05"), message.LicenseNumber)

		vehicleState := models.VehicleState{
			LicenseNumber:  message.LicenseNumber,
			EntryTimestamp: message.Timestamp,
		}

		err := h.repository.Save(vehicleState)
		if err != nil {
			return c.NoContent(http.StatusInternalServerError)
		}

		return c.NoContent(http.StatusOK)
	}
}
func (h *trafficControlHandlers) VehicleExit() echo.HandlerFunc {
	return func(c echo.Context) error {
		message := &events.VehicleRegistered{}

		if err := c.Bind(message); err != nil {
			h.logger.Error(err)
			return c.NoContent(http.StatusBadRequest)
		}

		if err := utils.ValidateStruct(c.Request().Context(), message); err != nil {
			h.logger.Error(err)
			return c.NoContent(http.StatusBadRequest)
		}

		h.logger.Infof(
			"EXIT detected in lane %d at %s of vehicle with license-number %s.",
			message.Lane, message.Timestamp.Format("15:04:05"), message.LicenseNumber)

		vehicleState, err := h.repository.Get(message.LicenseNumber)

		if err != nil {
			h.logger.Error(err)

			return c.NoContent(http.StatusNotFound)
		}

		vehicleState.ExitTimestamp = message.Timestamp

		err = h.repository.Save(vehicleState)
		if err != nil {
			h.logger.Error(err)
		}

		violation := h.service.DetermineSpeedingViolationInKmh(
			vehicleState.EntryTimestamp, vehicleState.ExitTimestamp)

		if violation > 0 {
			h.logger.Infof(
				"Speeding violation detected (%d KMh) of vehicle with license-number %s.",
				violation,
				vehicleState.LicenseNumber)

			speedingViolation := models.SpeedingViolation{
				VehicleId:      message.LicenseNumber,
				RoadId:         h.service.GetRoadId(),
				ViolationInKmh: violation,
				Timestamp:      message.Timestamp,
			}

			bodyBytes := new(bytes.Buffer)
			json.NewEncoder(bodyBytes).Encode(speedingViolation)

			client, err := dapr.NewClient()
			if err != nil {
				h.logger.Error(err)
				return c.NoContent(http.StatusInternalServerError)
			}

			ctx := context.Background()
			if err := client.PublishEvent(ctx, "pubsub", "speedingviolations", bodyBytes.Bytes()); err != nil {
				h.logger.Error(err)
				return c.NoContent(http.StatusInternalServerError)
			}
		}

		return c.NoContent(http.StatusOK)
	}
}
