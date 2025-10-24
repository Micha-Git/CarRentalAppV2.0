package mappers

import (
	"rentalmanagement/infrastructure/database/entities"
	"rentalmanagement/logic/model"
)

func ConvertRentableCarToRentableCarPersistenceEntity(rentableCar model.RentableCar) entities.RentableCarPersistenceEntity {
	return entities.RentableCarPersistenceEntity{
		Vin:         rentableCar.Vin.Vin,
		Brand:       rentableCar.Brand,
		Model:       rentableCar.Model,
		Location:    rentableCar.Location,
		PricePerDay: rentableCar.PricePerDay,
	}
}

func ConvertRentableCarPersistenceEntityToRentableCar(rentableCarPers entities.RentableCarPersistenceEntity) model.RentableCar {
	return model.RentableCar{
		Vin:         model.Vin{Vin: rentableCarPers.Vin},
		Brand:       rentableCarPers.Brand,
		Model:       rentableCarPers.Model,
		Location:    rentableCarPers.Location,
		PricePerDay: rentableCarPers.PricePerDay,
	}
}
