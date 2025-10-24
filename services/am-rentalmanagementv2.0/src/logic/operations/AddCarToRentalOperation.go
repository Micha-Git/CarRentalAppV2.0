package operations

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"rentalmanagement/logic/model"
)

var knownLimousines = []string{"Tesla Model 3", "Seat Leon"}
var knownCoupes = []string{"Fiat 500e", "Audi A3", "VW ID.2"}

func (ops RentableCarsCollectionOperations) AddCarToRental(vin model.Vin, location string) (model.RentableCar, error) {
	var msg string

	// Get car with provided vin if it exists
	car, err := ops.carRepository.GetCar(vin)
	if err != nil {
		msg = fmt.Sprintf("Failed to get information about car with VIN %s", vin.Vin)
		log.Error(msg, ": ", err)
		return model.RentableCar{}, fmt.Errorf("%s: %w", msg, err)
	}

	// Determine price per day to rent car with given vin
	// ToDo: determine vehicle type and therefore price per day based on the vin
	carName := fmt.Sprintf("%s %s", car.Brand, car.Model)
	pricePerDay := float32(20)
	if isStringInSlice(carName, knownLimousines) {
		pricePerDay = 50
	} else if isStringInSlice(carName, knownCoupes) {
		pricePerDay = 20
	} else if rand.Intn(2) == 0 {
		pricePerDay = 50
	}

	// Create rentable car
	rentableCar := model.RentableCar{
		Vin:         car.Vin,
		Brand:       car.Brand,
		Model:       car.Model,
		Location:    location,
		PricePerDay: pricePerDay,
	}

	// Add rentableCar to rentalRepository
	rentableCar, err = ops.rentableCarRepository.AddRentableCar(rentableCar)
	if err != nil {
		msg = fmt.Sprintf("Failed to add rentable car with VIN %s", vin.Vin)
		log.Error(msg, ": ", err)
		return model.RentableCar{}, fmt.Errorf("%s: %w", msg, err)
	}
	return rentableCar, nil
}

func isStringInSlice(str string, list []string) bool {
	for _, element := range list {
		if element == str {
			return true
		}
	}
	return false
}
