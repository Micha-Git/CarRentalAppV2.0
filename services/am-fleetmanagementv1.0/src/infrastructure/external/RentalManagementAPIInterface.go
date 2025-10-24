package external

import (
	"fleetmanagement/logic/model"
)

// .
type RentalManagementAPIInterface interface {
	AddCarToRental(vin model.Vin, location string) (model.RentableCar, error)
	RemoveRentableCar(vin model.Vin) (bool, error)
}
