package operations

import (
	"fmt"
	"rentalmanagement/logic/model"
	"time"

	log "github.com/sirupsen/logrus"
)

func (ops CustomerOperations) RentCar(customerId string, start time.Time, end time.Time, vin model.Vin) (model.Rental, error) {
	var msg string

	// Validate if StartDate is before EndDate
	if start.After(end) || start.Equal(end) {
		msg = "StartDate must be before EndDate"
		log.Warn(msg)
		return model.Rental{}, fmt.Errorf(msg)
	}

	// Check if the car exists in the car repository
	car, err := ops.rentableCarRepository.GetRentableCar(vin)

	if err != nil {
		msg = fmt.Sprintf("Car with VIN %s does not exist or error occured checking car existence", vin)
		log.Error(msg, ": ", err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	// Check if the car is available for the specified time range
	isCarAvailable, err := ops.rentalRepository.IsCarAvailableForRental(vin.Vin, start, end)
	if err != nil {
		msg = "Error checking car availability"
		log.Error(msg, ": ", err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	if !isCarAvailable {
		msg = fmt.Sprintf("Car with VIN %s is not available for the specified time range", vin)
		log.Warn(msg)
		return model.Rental{}, fmt.Errorf(msg)
	}

	days := int(end.Sub(start).Hours() / 24)
	totalPrice := float32(days) * car.PricePerDay

	rental := model.Rental{
		StartDate:  start,
		EndDate:    end,
		Car:        car,
		Price:      totalPrice,
		CustomerId: customerId,
	}

	// Add the rental
	rental, err = ops.rentalRepository.AddRental(rental)
	if err != nil {
		msg = "Error adding rental"
		log.Error(msg, ": ", err)
		return model.Rental{}, fmt.Errorf("%s: %w", msg, err)
	}

	return rental, nil
}
