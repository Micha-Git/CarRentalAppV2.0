package operations

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"rentalmanagement/logic/model"
)

func (ops RentalsCollectionOperations) ListCarRentals(vin model.Vin) ([]model.Rental, error) {
	var msg string

	if len(vin.Vin) == 0 {
		msg = "VIN can not be empty"
		log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Fetch all rentals of the car identified by the given vin
	rentals, err := ops.rentalRepository.ListRentalsByVin(vin)
	if err != nil {
		msg = fmt.Sprintf("Failed to list all rentals for car with VIN %s", vin.Vin)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentals, nil
}
