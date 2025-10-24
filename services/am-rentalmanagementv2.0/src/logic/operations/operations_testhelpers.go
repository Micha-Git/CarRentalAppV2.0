package operations

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"rentalmanagement/infrastructure/database"
	"rentalmanagement/logic/model"
	"testing"
)

var initialCarsInDatabase = []model.RentableCar{
	{
		Vin:         model.Vin{Vin: "JH4DB1561NS000567"},
		Brand:       "Tesla",
		Model:       "Model 3",
		Location:    "Karlsruhe",
		PricePerDay: 50,
	}, {
		Vin:         model.Vin{Vin: "JH4DB1561NS000568"},
		Brand:       "Seat",
		Model:       "Leon",
		Location:    "Karlsruhe",
		PricePerDay: 50,
	},
	{
		Vin:         model.Vin{Vin: "JH4DB1561NS000569"},
		Brand:       "Fiat",
		Model:       "500e",
		Location:    "Karlsruhe",
		PricePerDay: 20,
	},
	{
		Vin:         model.Vin{Vin: "JH4DB1561NS000570"},
		Brand:       "Audi",
		Model:       "A3",
		Location:    "Karlsruhe",
		PricePerDay: 20,
	},
	{
		Vin:         model.Vin{Vin: "JH4DB1561NS000571"},
		Brand:       "VW",
		Model:       "ID.2",
		Location:    "Mannheim",
		PricePerDay: 20,
	},
}

var carRepoCars = []model.Car{
	{
		Vin:   model.Vin{Vin: "JH4DB1561NS000667"},
		Brand: "Tesla",
		Model: "Model 3",
	},
	{
		Vin:   model.Vin{Vin: "JH4DB1561NS000668"},
		Brand: "Seat",
		Model: "Leon",
	},
}

var rentalRepository model.RentalRepositoryInterface
var rentableCarRepository model.RentableCarRepositoryInterface
var carRepositoryMockup model.CarRepositoryInterface
var rentalsCollectionOperations RentalsCollectionOperations
var customerOperations CustomerOperations
var rentableCarsCollectionOperations RentableCarsCollectionOperations

// SetupTestCase builds the basic scenario for every test of operations
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Logf("Setting up test case: '%s'", t.Name())
	var err error

	// Create SQLite in-memory database
	database.DB, err = gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models to database tables
	err = database.Migrate()
	if err != nil {
		t.Fatalf("Failed to migrate tables: %v", err)
	}

	// Create new repositories
	postgresRepository := database.NewPostgresRepository(database.DB)
	rentalRepository = &postgresRepository
	rentableCarRepository = &postgresRepository
	carRepositoryMockup = NewCarRepositoryMockup(carRepoCars)

	// Add rentable cars to DB
	for _, car := range initialCarsInDatabase {
		_, err := rentableCarRepository.AddRentableCar(car)
		if err != nil {
			t.Fatalf("Failed to store rentable car")
		}
	}

	// Create operations
	customerOperations = NewCustomerOperations(rentalRepository, rentableCarRepository)
	rentalsCollectionOperations = NewRentalsCollectionOperations(rentalRepository, rentableCarRepository)
	rentableCarsCollectionOperations = NewRentableCarsCollectionOperations(rentalRepository, rentableCarRepository, carRepositoryMockup)

	return func(t *testing.T) {
		t.Logf("Tearing down test case: '%s'", t.Name())
		err = database.CloseDB()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}
}

func carsHaveExpectedVins(expectedVins []model.Vin, receivedCars []model.RentableCar) bool {
	if len(expectedVins) != len(receivedCars) {
		return false
	}

	expectedVinsMap := make(map[string]bool)
	for _, vin := range expectedVins {
		expectedVinsMap[vin.Vin] = true
	}

	for _, car := range receivedCars {
		if !expectedVinsMap[car.Vin.Vin] {
			return false
		}
	}

	return true
}

func rentalArraysEqual(expectedRentals []model.Rental, receivedRentals []model.Rental) bool {
	if len(expectedRentals) != len(receivedRentals) {
		return false
	}

	expectedRentalsMap := make(map[string]model.Rental)
	for _, rental := range expectedRentals {
		expectedRentalsMap[rental.Id] = rental
	}

	for _, receivedRental := range receivedRentals {
		expectedRental, receivedRentalIdKeyInMap := expectedRentalsMap[receivedRental.Id]
		if !receivedRentalIdKeyInMap || expectedRental != receivedRental {
			return false
		}
	}

	return true
}
