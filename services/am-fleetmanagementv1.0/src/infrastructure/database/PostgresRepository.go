package database

import (
	"fleetmanagement/infrastructure/database/entities"
	"fleetmanagement/infrastructure/database/mappers"
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	DB *gorm.DB
}

// NewMainRepository returns a new MainRepository instance
func NewPostgresRepository(db *gorm.DB) *PostgresRepository {
	return &PostgresRepository{
		DB: db,
	}
}

// Car Methods

// adds car to fleet
func (repo *PostgresRepository) AddCarToFleet(car model.Car, fleetId string, location string) (model.Car, error) {
	var carResult model.Car
	carPers := mappers.ConvertCarToCarPersistenceEntity(car, fleetId, location)

	if err := repo.DB.Create(&carPers).Error; err != nil {
		msg := fmt.Sprintf("Database Failed to add car VIN %s", car.Vin.Vin)
		log.Error(msg, ": ", err)
		return carResult, fmt.Errorf("%s: %w", msg, err)
	}

	carResult = car

	return carResult, nil

}

// returns a car given the vin as string
func (repo *PostgresRepository) GetCar(vin model.Vin) (model.Car, error) {
	//TODO implement accordingly, this is just the DoesRentalExist function from AM-RentalManagement
	var msg string
	var carPers entities.CarPersistenceEntity
	var car model.Car

	// Check if the car exists
	if err := repo.DB.Where("vin = ?", vin.Vin).First(&carPers).Error; err != nil {
		msg = fmt.Sprintf("Database failed to find car with vin %s", vin.Vin)
		log.Error(msg, ": ", err)
		return car, fmt.Errorf("%s: %w", msg, err)
	}

	carObject := mappers.ConvertCarPersistenceEntityToCar(carPers)
	car = carObject
	// If we get to this point, it means a rental with the given ID exists
	return car, nil
}

// removes a car from a fleet given its vin as Vin
func (repo *PostgresRepository) RemoveCar(vin model.Vin) (bool, error) {
	var msg string
	var firstEntry entities.CarPersistenceEntity

	// Directly delete the fleet based on the ID
	if err := repo.DB.Where("vin = ?", vin.Vin).First(&firstEntry).Error; err != nil {
		msg = fmt.Sprintf("Database failed to find car with vin %s", vin.Vin)
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	} else {
		if err := repo.DB.Delete(&firstEntry).Error; err != nil {
			msg = fmt.Sprintf("Database failed to delete car with vin %s", vin.Vin)
			log.Error(msg, ": ", err)
			return false, fmt.Errorf("%s: %w", msg, err)
		}
	}

	return true, nil
}

func (repo *PostgresRepository) ListAllCars(fleetId string) ([]model.Car, error) {
	var carPersEntities []entities.CarPersistenceEntity
	var cars []model.Car

	// Query for Fleets by fleetID
	if err := repo.DB.Where("fleet_Id = ?", fleetId).Find(&carPersEntities).Error; err != nil {
		msg := "Database failed to list cars"
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Convert persistence entities to domain entities
	for _, carPers := range carPersEntities {
		car := mappers.ConvertCarPersistenceEntityToCar(carPers)
		cars = append(cars, car)
	}

	return cars, nil
}

// Fleet Methods

// GetFleet returns a Fleet from the database using the Fleet's ID
func (repo *PostgresRepository) GetFleet(fleetID string) (model.Fleet, error) {
	var msg string
	var fleetPersEntity entities.FleetPersistenceEntity
	var fleet model.Fleet

	// Get the fleet based on the ID
	if err := repo.DB.Where("fleet_id = ?", fleetID).First(&fleetPersEntity).Error; err != nil {
		msg = fmt.Sprintf("Database failed to get fleet with ID %s", fleetID)
		log.Error(msg, ": ", err)
		return fleet, fmt.Errorf("%s: %w", msg, err)
	}

	// Get fleet's cars
	cars, err := repo.ListAllCars(fleetID)
	if err != nil {
		msg = fmt.Sprintf("Failed to get cars for fleet with ID %s", fleetID)
		log.Error(msg, ": ", err)
		return fleet, fmt.Errorf("%s: %w", msg, err)
	}

	// Convert persistence entity to domain entity
	fleetObj := mappers.ConvertFleetPersistenceEntityToFleet(fleetPersEntity) // Assuming you have a function like this
	fleet = fleetObj
	fleet.Cars = cars

	return fleetObj, nil
}

// AddFleet adds a fleet to the database
func (repo *PostgresRepository) AddFleet(fleet model.Fleet) (model.Fleet, error) {
	var msg string
	fleetPres := mappers.ConvertFleetToFleetPersistenceEntity(fleet)

	if err := repo.DB.Create(&fleetPres).Error; err != nil {
		msg = fmt.Sprintf("Database Failed to add fleet with fleetID %s", fleet.FleetId)
		log.Error(msg, ": ", err)
		return fleet, fmt.Errorf("%s: %w", msg, err)
	}

	return mappers.ConvertFleetPersistenceEntityToFleet(fleetPres), nil
}
