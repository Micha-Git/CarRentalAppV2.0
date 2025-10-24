package operations

import (
	model "rentalmanagement/logic/model"
	"time"
)

type RentalsCollectionOperationsInterface interface {
	ListAvailableCars(startDate, endDate time.Time, location string) ([]model.RentableCar, error)
	ListCarRentals(vin model.Vin) ([]model.Rental, error)
	ListCustomerRentals(customerId string) ([]model.Rental, error)
}
