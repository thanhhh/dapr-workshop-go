package repositories

import (
	"dapr-workshop-go/vehicle-registration-service/internal/models"
	vr "dapr-workshop-go/vehicle-registration-service/internal/vehicle_registration"
	"math/rand"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type inmemoryRepository struct {
	vehicles map[string]models.VehicleInfo
}

var (
	vehicleBrands = [...]string{
		"Mercedes", "Toyota", "Audi", "Volkswagen", "Seat", "Renault", "Skoda",
		"Kia", "Citroën", "Suzuki", "Mitsubishi", "Fiat", "Opel"}

	vehicleModels = map[string][]string{
		"Mercedes":   []string{"A Class", "B Class", "C Class", "E Class", "SLS", "SLK"},
		"Toyota":     []string{"Yaris", "Avensis", "Rav 4", "Prius", "Celica"},
		"Audi":       []string{"A3", "A4", "A6", "A8", "Q5", "Q7"},
		"Volkswagen": []string{"Golf", "Pasat", "Tiguan", "Caddy"},
		"Seat":       []string{"Leon", "Arona", "Ibiza", "Alhambra"},
		"Renault":    []string{"Megane", "Clio", "Twingo", "Scenic", "Captur"},
		"Skoda":      []string{"Octavia", "Fabia", "Superb", "Karoq", "Kodiaq"},
		"Kia":        []string{"Picanto", "Rio", "Ceed", "XCeed", "Niro", "Sportage"},
		"Citroën":    []string{"C1", "C2", "C3", "C4", "C4 Cactus", "Berlingo"},
		"Suzuki":     []string{"Ignis", "Swift", "Vitara", "S-Cross", "Swace", "Jimny"},
		"Mitsubishi": []string{"Space Star", "ASX", "Eclipse Cross", "Outlander PHEV"},
		"Ford":       []string{"Focus", "Ka", "C-Max", "Fusion", "Fiesta", "Mondeo", "Kuga"},
		"BMW":        []string{"1 Serie", "2 Serie", "3 Serie", "5 Serie", "7 Serie", "X5"},
		"Fiat":       []string{"500", "Panda", "Punto", "Tipo", "Multipla"},
		"Opel":       []string{"Karl", "Corsa", "Astra", "Crossland X", "Insignia"},
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
