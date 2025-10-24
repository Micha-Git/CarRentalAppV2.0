package operations

import (
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops FleetOperations) AddCarToFleet(vin model.Vin, fleetId string, location string) (model.Car, error) {
	var msg string

	//Check if fleet exists.
	_, err := ops.repository.GetFleet(fleetId)
	if err != nil {
		msg = fmt.Sprintf("Fleet with ID %s does not exist", fleetId)
		log.Error(msg, err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	//Check if vin already present.
	_, err = ops.repository.GetCar(vin)
	if err == nil {
		msg = fmt.Sprintf("Car with vin %s already present in the database", vin)
		log.Error(msg, nil)
		return model.Car{}, fmt.Errorf("%s", msg)
	}

	car, err := ops.carApi.GetCar(vin)
	if err != nil {
		msg = fmt.Sprintf("Failed to get information about car with VIN %s", vin.Vin)
		log.Error(msg, err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	//Add car to fleet.
	addCar, err := ops.repository.AddCarToFleet(car, fleetId, location)
	if err != nil {
		msg = fmt.Sprintf("Failed to add car with vin %s", vin)
		log.Error(msg, err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	//rpc call AM-RentalManagement
	resp, err := ops.rentalManagementApi.AddCarToRental(vin, location)
	if err != nil || (resp == model.RentableCar{}) {
		msg = fmt.Sprintf("Failed to add car with vin %s and location %s", vin, location)
		log.Error(msg, err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	return addCar, err

}
