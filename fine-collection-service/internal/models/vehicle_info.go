package models

type VehicleInfo struct {
	VehicleId  string `json:"vehicle_id" validate:"required"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	OwnerName  string `json:"owner_name" validate:"required"`
	OwnerEmail string `json:"owner_email" validate:"required"`
}
