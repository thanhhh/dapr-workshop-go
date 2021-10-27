package repositories

import (
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"

	"dapr-workshop-go/vehicle-registration-service/internal/models"
	vr "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration"
)

type inmemoryRepository struct {
	vehicles map[string]models.VehicleInfo
}

var (
	vehicleBrands = [...]string{
		"Mercedes", "Toyota", "Audi", "Volkswagen", "Seat", "Renault", "Skoda",
		"Kia", "Citroën", "Suzuki", "Mitsubishi", "Fiat", "Opel"}

	vehicleModels = map[string][]string{
		"Mercedes":   {"A Class", "B Class", "C Class", "E Class", "SLS", "SLK"},
		"Toyota":     {"Yaris", "Avensis", "Rav 4", "Prius", "Celica"},
		"Audi":       {"A3", "A4", "A6", "A8", "Q5", "Q7"},
		"Volkswagen": {"Golf", "Pasat", "Tiguan", "Caddy"},
		"Seat":       {"Leon", "Arona", "Ibiza", "Alhambra"},
		"Renault":    {"Megane", "Clio", "Twingo", "Scenic", "Captur"},
		"Skoda":      {"Octavia", "Fabia", "Superb", "Karoq", "Kodiaq"},
		"Kia":        {"Picanto", "Rio", "Ceed", "XCeed", "Niro", "Sportage"},
		"Citroën":    {"C1", "C2", "C3", "C4", "C4 Cactus", "Berlingo"},
		"Suzuki":     {"Ignis", "Swift", "Vitara", "S-Cross", "Swace", "Jimny"},
		"Mitsubishi": {"Space Star", "ASX", "Eclipse Cross", "Outlander PHEV"},
		"Ford":       {"Focus", "Ka", "C-Max", "Fusion", "Fiesta", "Mondeo", "Kuga"},
		"BMW":        {"1 Serie", "2 Serie", "3 Serie", "5 Serie", "7 Serie", "X5"},
		"Fiat":       {"500", "Panda", "Punto", "Tipo", "Multipla"},
		"Opel":       {"Karl", "Corsa", "Astra", "Crossland X", "Insignia"},
	}
)

func NewInMemoryRepository() vr.VehicleInfoRepository {
	return &inmemoryRepository{vehicles: make(map[string]models.VehicleInfo)}
}

func GetRandNumber(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

func GetRandomBrand() string {
	return vehicleBrands[GetRandNumber(len(vehicleBrands))]
}
func GetRandomModel(brand string) string {
	models := vehicleModels[brand]
	return models[GetRandNumber(len(models))]
}

func (r *inmemoryRepository) Get(licenseNumber string) models.VehicleInfo {
	brand := GetRandomBrand()
	model := GetRandomModel(brand)
	ownerName := gofakeit.Name()
	ownerEmail := strings.ReplaceAll(ownerName, " ", ".") + "@outlook.com"

	return models.VehicleInfo{
		VehicleId:  licenseNumber,
		Brand:      brand,
		Model:      model,
		OwnerName:  ownerName,
		OwnerEmail: ownerEmail,
	}
}
