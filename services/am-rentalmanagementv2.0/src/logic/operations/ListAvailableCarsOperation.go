package operations

import (
	"fmt"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
)

func (ops RentalsCollectionOperations) ListAvailableCars(startDate, endDate time.Time, location string) ([]model.RentableCar, error) {
	var msg string

	// Validate if StartDate is before EndDate
	if startDate.After(endDate) || startDate.Equal(endDate) {
		msg = "StartDate must be before EndDate"
		log.Warn(msg)
		return nil, fmt.Errorf(msg)
	}

	// Fetch all cars; ideally, this would be optimized to only fetch available cars from database by querying
	rentableCars, err := ops.rentableCarRepository.ListRentableCarsByLocation(location)
	if err != nil {
		msg = "Failed to list all available cars"
		log.Error(msg, ": ", err)
		return nil, fmt.Errorf("%s: %w", msg, err)
	}

	// Pre-allocate the slice to avoid dynamic resizing
	availableCars := make([]model.RentableCar, 0, len(rentableCars))

	for _, rentableCar := range rentableCars {
		isAvailable, err := ops.rentalRepository.IsCarAvailableForRental(rentableCar.Vin.Vin, startDate, endDate)
		if err != nil {
			msg = fmt.Sprintf("Failed to check availability for car with VIN %s", rentableCar.Vin.Vin)
			log.Error(msg, ": ", err)
			return nil, fmt.Errorf("%s: %w", msg, err)
		}

		if isAvailable {
			availableCars = append(availableCars, rentableCar)
		}
	}

	return availableCars, nil
}
