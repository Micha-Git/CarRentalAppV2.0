package operations

import (
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops FleetOperations) RemoveCarFromFleet(vin model.Vin) (bool, error) {
	var msg string

	//Check if car exists.
	_, err := ops.repository.GetCar(vin)
	if err != nil {
		msg = fmt.Sprintf("Car with ID %s does not exist", vin)
		log.Error(msg, err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	//Remove rentable Car through a rpc call
	resp, err := ops.rentalManagementApi.RemoveRentableCar(vin)
	if err != nil || !resp {
		msg = fmt.Sprintf("Failed to remove rentable car with vin %s", vin)
		log.Error(msg, err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	//Delete Car from given fleet.
	removal, err := ops.repository.RemoveCar(vin)
	if err != nil || !removal {
		msg = fmt.Sprintf("Failed to remove the car with ID %s", vin)
		log.Error(msg, err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	return true, err
}
