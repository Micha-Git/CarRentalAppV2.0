package operations

import (
	model "rentalmanagement/logic/model"
)

type RentableCarsCollectionOperationsInterface interface {
	AddCarToRental(vin model.Vin, location string) (model.RentableCar, error)
	RemoveRentableCar(vin model.Vin) error
}
