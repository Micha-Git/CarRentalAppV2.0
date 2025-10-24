package database

import (
	"fmt"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/infrastructure/database/mappers"
	"rentalmanagement/logic/model"
	"time"
)

// PostgresRepository implements the PostgresRepositoryInterface
type PostgresRepository struct {
	DB *gorm.DB
}

// NewPostgresRepository returns a new PostgresRepository instance
func NewPostgresRepository(db *gorm.DB) PostgresRepository {
	return PostgresRepository{
		DB: db,
	}
}

func (repo *PostgresRepository) AddRentableCar(rentableCar model.RentableCar) (model.RentableCar, error) {
	rentableCarPers := mappers.ConvertRentableCarToRentableCarPersistenceEntity(rentableCar)

	if err := repo.DB.Create(&rentableCarPers).Error; err != nil {
		msg := fmt.Sprintf("Database failed to add rentable car with VIN %s", rentableCar.Vin.Vin)
		log.Error(msg, ": ", err)
		return rentableCar, fmt.Errorf("%s: %w", msg, err)
	}

	return mappers.ConvertRentableCarPersistenceEntityToRentableCar(rentableCarPers), nil
}

func (repo *PostgresRepository) GetRentableCar(vin model.Vin) (model.RentableCar, error) {
	var rentableCarPers entities.RentableCarPersistenceEntity

	// Query for rentable car by vin
	if err := repo.DB.Where("vin = ?", vin.Vin).First(&rentableCarPers).Error; err != nil {
		msg := fmt.Sprintf("Database failed to find rentable car with VIN %s", vin)
		log.Error(msg, ": ", err)
		return model.RentableCar{}, fmt.Errorf("%s: %w", msg, err)
	}

	return mappers.ConvertRentableCarPersistenceEntityToRentableCar(rentableCarPers), nil
}

func (repo *PostgresRepository) ListRentableCarsByLocation(location string) ([]model.RentableCar, error) {
	rentableCars, err := repo.listRentableCarsByCondition("location = ?", location)
	if err != nil {
		msg := fmt.Sprintf("Databse failed to list rentable cars by location %s", location)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentableCars, nil
}

func (repo *PostgresRepository) listRentableCarsByCondition(query string, args ...interface{}) ([]model.RentableCar, error) {
	var rentableCarPersEntities []entities.RentableCarPersistenceEntity
	var rentableCars []model.RentableCar

	// Query for rentable cars matching the given condition
	if err := repo.DB.Where(query, args...).Find(&rentableCarPersEntities).Error; err != nil {
		return nil, err
	}

	// Convert persistence entities to domain entities
	for _, rentableCarPers := range rentableCarPersEntities {
		rentableCar := mappers.ConvertRentableCarPersistenceEntityToRentableCar(rentableCarPers)
		rentableCars = append(rentableCars, rentableCar)
	}

	return rentableCars, nil
}

func (repo *PostgresRepository) RemoveRentableCar(vin string) error {
	var msg string

	tx := repo.DB.Where("vin = ?", vin).Delete(&entities.RentableCarPersistenceEntity{})
	err := tx.Error
	if err == nil && tx.RowsAffected == 0 {
		err = fmt.Errorf("There is no rentable car matching the given condition")
	}
	if err != nil {
		msg = fmt.Sprintf("Database failed to remove rentable car with VIN %s", vin)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}

// AddRental adds a new rental to the database
func (repo *PostgresRepository) AddRental(rental model.Rental) (model.Rental, error) {
	rental.Id = uuid.New().String()
	rentalPers := mappers.ConvertRentalToRentalPersistenceEntity(rental)

	err := repo.DB.Create(&rentalPers).Error
	if err != nil {
		msg := fmt.Sprintf("Database failed to create rental for car with VIN %s", rental.Car.Vin.Vin)
		log.Error(msg, ": ", err)
		return rental, fmt.Errorf("%s: %w", msg, err)
	}

	return mappers.ConvertRentalPersistenceEntityToRental(rentalPers, rental.Car), nil
}

// IsCarAvailableForRental checks if a car with the given VIN is available for rental between startDate and endDate
func (repo *PostgresRepository) IsCarAvailableForRental(vin string, startDate, endDate time.Time) (bool, error) {
	var count int64

	err := repo.DB.Model(&entities.RentalPersistenceEntity{}).
		Where("vin = ?", vin).
		Where("(start_date <= ?) AND (end_date >= ?)", endDate, startDate).
		Count(&count).Error
	if err != nil {
		msg := fmt.Sprintf("Database failed to check availability for car with VIN %s", vin)
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	if count > 0 {
		log.Warn(fmt.Sprintf("Car with VIN %s is not available for the given date range", vin))
		return false, nil
	}

	return true, nil
}

func (repo *PostgresRepository) ListRentalsByCustomerId(customerId string) ([]model.Rental, error) {
	rentals, err := repo.listRentalsByCondition("customer_id = ?", customerId)
	if err != nil {
		msg := fmt.Sprintf("Database failed to list rentals for customer with ID %s", customerId)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentals, nil
}

func (repo *PostgresRepository) ListRentalsByVin(vin model.Vin) ([]model.Rental, error) {
	rentals, err := repo.listRentalsByCondition("vin = ?", vin.Vin)
	if err != nil {
		msg := fmt.Sprintf("Database failed to list rentals for cars with VIN %s", vin.Vin)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentals, nil
}

func (repo *PostgresRepository) listRentalsByCondition(query string, args ...interface{}) ([]model.Rental, error) {
	var rentalPersEntities []entities.RentalPersistenceEntity
	var rentals []model.Rental

	// Query for rentals matching the given condition
	err := repo.DB.Where(query, args...).Find(&rentalPersEntities).Error
	if err != nil {
		return nil, err
	}

	// Convert persistence entities to domain entities
	for _, rentalPers := range rentalPersEntities {
		rentableCar, err := repo.GetRentableCar(model.Vin{Vin: rentalPers.Vin})

		if err != nil {
			msg := fmt.Sprintf("Failed to convert database rental %v to rental object", rentalPers)
			log.Error(msg, ":", err)
		}

		rental := mappers.ConvertRentalPersistenceEntityToRental(rentalPers, rentableCar)
		rentals = append(rentals, rental)
	}

	return rentals, nil
}

func (repo *PostgresRepository) DeleteRental(rentalId string) error {
	err := repo.deleteRentalByCondition("id = ?", rentalId)
	if err != nil {
		msg := fmt.Sprintf("Database failed to delete rental with ID %s", rentalId)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}

func (repo *PostgresRepository) DeleteRentalOfCustomer(rentalId string, customerId string) error {
	err := repo.deleteRentalByCondition("id = ? AND customer_id = ?", rentalId, customerId)
	if err != nil {
		msg := fmt.Sprintf("Database failed to delete rental with ID %s belonging to customer with ID %s", rentalId, customerId)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}

func (repo *PostgresRepository) deleteRentalByCondition(query string, args ...interface{}) error {
	tx := repo.DB.Where(query, args...).Delete(&entities.RentalPersistenceEntity{})

	err := tx.Error
	if err == nil && tx.RowsAffected == 0 {
		err = fmt.Errorf("There is no rental matching the given condition")
	}

	return err
}
