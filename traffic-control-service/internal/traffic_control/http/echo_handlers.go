package http

import (
	"bytes"
	"dapr-workshop-go/traffic-control-service/config"
	"dapr-workshop-go/traffic-control-service/internal/events"
	"dapr-workshop-go/traffic-control-service/internal/models"
	tc "dapr-workshop-go/traffic-control-service/internal/traffic_control"
	"dapr-workshop-go/traffic-control-service/pkg/logger"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

type trafficControlHandlers struct {
	cfg        *config.Config
	service    tc.SpeedingViolationCalculator
	repository tc.VehicleStateRepository
	logger     logger.Logger
}

func NewTrafficControlHandlers(cfg *config.Config, service tc.SpeedingViolationCalculator, repository tc.VehicleStateRepository, logger logger.Logger) tc.Handlers {
	return &trafficControlHandlers{cfg: cfg, service: service, repository: repository, logger: logger}
}

func (h *trafficControlHandlers) VehicleEntry() echo.HandlerFunc {
	return func(c echo.Context) error {
		message := &events.VehicleRegistered{}

		if err := c.Bind(message); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

		h.logger.Info(
			"ENTRY detected in lane %d at %s of vehicle with license-number {msg.LicenseNumber}.",
			message.Lane, message.Timestamp.Format("hh:mm:ss"))

		vehicleState := models.VehicleState{
			LicenseNumber:  message.LicenseNumber,
			EntryTimestamp: message.Timestamp,
		}

		h.repository.Save(vehicleState)

		return c.NoContent(http.StatusOK)
	}
}
func (h *trafficControlHandlers) VehicleExit() echo.HandlerFunc {
	return func(c echo.Context) error {
		message := &events.VehicleRegistered{}

		if err := c.Bind(message); err != nil {
			return c.NoContent(http.StatusBadRequest)
		}

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
			h.logger.Info(
				"Speeding violation detected (%d KMh) of vehicle with license-number %s.",
				violation,
				vehicleState.LicenseNumber)

			speedingViolation := models.SpeedingViolation{
				LicenseNumber:  message.LicenseNumber,
				RoadId:         h.service.GetRoadId(),
				ViolationInKmh: violation,
				Timestamp:      message.Timestamp,
			}

			data, err := json.Marshal(speedingViolation)
			if err != nil {
				h.logger.Error(err)
				return c.NoContent(http.StatusInternalServerError)
			}
			req, err := http.NewRequest("POST", "http://localhost:6001/collectfine", bytes.NewBuffer(data))
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				h.logger.DPanic(err)

				return c.NoContent(http.StatusInternalServerError)
			}
			defer resp.Body.Close()
		}

		return c.NoContent(http.StatusOK)
	}
}
