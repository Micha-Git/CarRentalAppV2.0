package operations

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func (ops CustomerOperations) CancelRental(customerId string, rentalId string) error {
	err := ops.rentalRepository.DeleteRentalOfCustomer(rentalId, customerId)
	if err != nil {
		msg := fmt.Sprintf("Failed to cancel the rental with ID %s", rentalId)
		log.Error(msg, ": ", err)
		return fmt.Errorf("%s: %w", msg, err)
	}

	return nil
}
