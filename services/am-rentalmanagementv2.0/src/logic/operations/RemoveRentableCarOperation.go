package operations

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"rentalmanagement/logic/model"
)

func (ops RentableCarsCollectionOperations) RemoveRentableCar(vin model.Vin) error {
	handleFailure := func(vin model.Vin, err error) error {
		msg := fmt.Sprintf("Failed to remove rentable car with VIN %s", vin.Vin)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	rentals, err := ops.rentalRepository.ListRentalsByVin(vin)
	if err != nil {
		return handleFailure(vin, err)
	}
	for _, rental := range rentals {
		err = ops.rentalRepository.DeleteRental(rental.Id)
		if err != nil {
			return handleFailure(vin, err)
		}
	}

	err = ops.rentableCarRepository.RemoveRentableCar(vin.Vin)
	if err != nil {
		return handleFailure(vin, err)
	}

	return nil
}
