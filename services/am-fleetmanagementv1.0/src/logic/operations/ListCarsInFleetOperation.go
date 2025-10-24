package operations

import (
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops FleetOperations) ListCarsInFleet(fleetId string) ([]model.Car, error) {
	var msg string

	//Check if fleet exists.
	_, err := ops.repository.GetFleet(fleetId)
	if err != nil {
		msg = fmt.Sprintf("Fleet with ID %s does not exist", fleetId)
		log.Error(msg, err)
		return []model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	//Retrieve all cars in given fleet.
	cars, err := ops.repository.ListAllCars(fleetId)
	if err != nil {
		msg = fmt.Sprintf("Failed to retrieve list of cars in fleet with ID %s", fleetId)
		log.Error(msg, err)
		return []model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	return cars, err

}
