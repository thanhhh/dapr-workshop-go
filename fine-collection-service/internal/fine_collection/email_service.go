package finecollection

import "dapr-workshop-go/fine-collection-service/internal/models"

type EmailService interface {
	SendMail(speedingViolation models.SpeedingViolation,
		vehicleInfo models.VehicleInfo,
		fine string) error
}
