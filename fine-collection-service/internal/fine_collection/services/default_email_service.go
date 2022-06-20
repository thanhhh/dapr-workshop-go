package services

import (
	finecollection "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
	"dapr-workshop-go/pkg/logger"
)

type defaultEmailService struct {
	logger logger.Logger
}

func NewEmailService(logger logger.Logger) finecollection.EmailService {
	return &defaultEmailService{
		logger: logger,
	}
}

func (e defaultEmailService) SendMail(speedingViolation models.SpeedingViolation,
	vehicleInfo models.VehicleInfo,
	fine string) error {

	return nil
}
