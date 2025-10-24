package operations

import (
	"fmt"
	"rentalmanagement/logic/model"

	log "github.com/sirupsen/logrus"
)

func (ops RentalsCollectionOperations) ListCustomerRentals(customerId string) ([]model.Rental, error) {
	var msg string

	if len(customerId) == 0 {
		msg = "Customer ID can't be empty"
		log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Fetch all rentals of the car identified by the given customer id
	rentals, err := ops.rentalRepository.ListRentalsByCustomerId(customerId)
	if err != nil {
		msg = fmt.Sprintf("Failed to list all rentals of customer with ID %s", customerId)
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	return rentals, nil
}
