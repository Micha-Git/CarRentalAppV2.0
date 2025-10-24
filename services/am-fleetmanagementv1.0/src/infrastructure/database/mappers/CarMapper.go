package mappers

import (
	"fleetmanagement/infrastructure/database/entities"
	"fleetmanagement/logic/model"
)

func ConvertCarToCarPersistenceEntity(car model.Car, fleetID string, location string) entities.CarPersistenceEntity {

	return entities.CarPersistenceEntity{
		Vin:      car.Vin.Vin,
		Model:    car.Model,
		Brand:    car.Brand,
		Location: location,
		FleetId:  fleetID,
	}
}

func ConvertCarPersistenceEntityToCar(carPers entities.CarPersistenceEntity) model.Car {
	var vin model.Vin
	vin.Vin = carPers.Vin

	return model.Car{
		Vin:   vin,
		Brand: carPers.Brand,
		Model: carPers.Model,
	}
}
