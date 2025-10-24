package operations

import (
	"rentalmanagement/logic/model"
	"time"
)

type CustomerOperationsInterface interface {
	RentCar(customerId string, start time.Time, end time.Time, vin model.Vin) (model.Rental, error)
	CancelRental(customerId string, rentalId string) error
}
