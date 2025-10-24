package operations

import (
	"fleetmanagement/infrastructure/database"
	"fleetmanagement/infrastructure/external"
	"fleetmanagement/logic/model"

	"testing"
)

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

var postgresRepo model.PostgresRepositoryInterface
var carOperations CarOperations
var fleetOperations FleetOperations

// SetupTestCase builds the basic scenario for every test of operations
func setupTestCase(t *testing.T) func(t *testing.T) {
	t.Logf("Setting up test case: '%s'", t.Name())

	// Create SQLite in-memory database
	_, err := database.ConnectDB(true)
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}

	// Migrate models to database tables
	err = database.Migrate()
	if err != nil {
		t.Fatalf("Failed to migrate tables: %v", err)
	}

	// Create new repository
	postgresRepo = database.NewPostgresRepository(database.InMemDB)

	addTestData(postgresRepo.(*database.PostgresRepository))

	// Create operations
	carAPI := external.NewCarAPIStub(carRepoCars)

	rentalManagementApi := external.NewRentalManagementAPIStub(carRepoCars)

	fleetOperations = NewFleetOperations(postgresRepo, carAPI, rentalManagementApi)
	carOperations = NewCarOperations(postgresRepo, carAPI)

	return func(t *testing.T) {
		t.Logf("Tearing down test case: '%s'", t.Name())
		err = database.CloseDB()
		if err != nil {
			t.Fatalf("Failed to close database: %v", err)
		}
	}
}

func addTestData(repo *database.PostgresRepository) {

	testFleet := model.Fleet{
		Cars:         []model.Car{carRepoCars[0]},
		Location:     "Karlsruhe",
		FleetManager: "Fred",
		FleetId:      "1",
	}

	repo.AddFleet(testFleet)

	repo.AddCarToFleet(carRepoCars[0], "1", "Karlsruhe")
}
