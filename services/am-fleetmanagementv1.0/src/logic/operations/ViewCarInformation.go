package operations

import (
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops CarOperations) ViewCarInformation(vin model.Vin) (model.Car, error) {
	var msg string

	//retrieve the car with the matching vin.
	car, err := ops.repository.GetCar(vin)
	if err != nil {
		msg = fmt.Sprintf("Failed to retrieve car with vin %s", vin)
		log.Error(msg, err)
		return model.Car{}, fmt.Errorf("%s: %w", msg, err)
	}

	return car, err

}
