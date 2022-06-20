package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"dapr-workshop-go/pkg/logger"

	finecollection "dapr-workshop-go/fine-collection-service/internal/fine_collection"
	"dapr-workshop-go/fine-collection-service/internal/models"
)

type defaultEmailService struct {
	logger logger.Logger
}

type MessageMetadata struct {
	Subject   string `json:"subject"`
	EmailTo   string `json:"emailTo"`
	EmailFrom string `json:"emailFrom"`
}
type MessageData struct {
	Data      string          `json:"data"`
	Operation string          `json:"operation"`
	Metadata  MessageMetadata `json:"metadata"`
}

type MailData struct {
	Now            time.Time
	OwnerName      string
	VehicleId      string
	Brand          string
	Model          string
	RoadId         string
	Timestamp      time.Time
	ViolationInKmh int
	Fine           string
}

func NewEmailService(logger logger.Logger) finecollection.EmailService {
	return &defaultEmailService{
		logger: logger,
	}
}

func (e defaultEmailService) SendMail(speedingViolation models.SpeedingViolation,
	vehicleInfo models.VehicleInfo,
	fine string) error {

	mailData := MailData{
		Now:            time.Now(),
		OwnerName:      vehicleInfo.OwnerName,
		VehicleId:      vehicleInfo.VehicleId,
		Brand:          vehicleInfo.Brand,
		Model:          vehicleInfo.Model,
		RoadId:         speedingViolation.RoadId,
		Timestamp:      speedingViolation.Timestamp,
		ViolationInKmh: speedingViolation.ViolationInKmh,
		Fine:           fine,
	}

	mailBody, err := ParseTemplate("templates/email.html", mailData)

	if err != nil {
		return err
	}

	messageData := MessageData{
		Data:      mailBody,
		Operation: "create",
		Metadata: MessageMetadata{
			Subject:   "Fine for exceeding the speed limit.",
			EmailTo:   vehicleInfo.OwnerEmail,
			EmailFrom: "test@domain.org",
		},
	}

	messageJson, err := json.Marshal(messageData)
	if err != nil {

		return fmt.Errorf("SendMail encode json error: %v", err)
	}

	resp, err := http.Post("http://localhost:3601/v1.0/bindings/sendmail", "application/json", bytes.NewBuffer(messageJson))

	if err != nil {
		return fmt.Errorf("SendMail create http request Dapr binding sendmail error: %v", err)
	}

	defer resp.Body.Close()

	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", fmt.Errorf("ParseTemplate: parse files error: %w", err)
	}

	buf := new(bytes.Buffer)

	if err = t.Execute(buf, data); err != nil {
		return "", fmt.Errorf("ParseTemplate: execute template error: %w", err)
	}

	return buf.String(), nil
}
