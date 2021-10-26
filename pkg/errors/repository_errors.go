package errors

import "errors"

var (
	ErrVehicleStateRecordNotFound  = errors.New("Vehicle state record not found.")
	ErrLicenseKeyInvalidOrNotFound = errors.New("Invalid or no license key specified.")
)
