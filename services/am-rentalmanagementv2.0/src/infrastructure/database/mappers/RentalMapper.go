package mappers

import (
	"github.com/google/uuid"
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/logic/model"
)

func ConvertRentalToRentalPersistenceEntity(rental model.Rental) entities.RentalPersistenceEntity {
	rentalUUID, err := uuid.Parse(rental.Id)
	if err != nil {
		panic("Could not convert PersistenceEntity")
	}
	return entities.RentalPersistenceEntity{
		Id:         rentalUUID,
		StartDate:  rental.StartDate,
		EndDate:    rental.EndDate,
		Vin:        rental.Car.Vin.Vin,
		Price:      rental.Price,
		CustomerId: rental.CustomerId,
	}
}

func ConvertRentalPersistenceEntityToRental(rentalPers entities.RentalPersistenceEntity, rentableCar model.RentableCar) model.Rental {
	rentalID := rentalPers.Id.String()
	return model.Rental{
		Id:         rentalID,
		StartDate:  rentalPers.StartDate,
		EndDate:    rentalPers.EndDate,
		Car:        rentableCar,
		Price:      rentalPers.Price,
		CustomerId: rentalPers.CustomerId,
	}
}
