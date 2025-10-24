package model

import "time"

type RentalRepositoryInterface interface {
	AddRental(rental Rental) (Rental, error)
	IsCarAvailableForRental(vin string, startDate, endDate time.Time) (bool, error)
	ListRentalsByCustomerId(customerId string) ([]Rental, error)
	ListRentalsByVin(vin Vin) ([]Rental, error)
	DeleteRental(rentalId string) error
	DeleteRentalOfCustomer(rentalId string, customerId string) error
}
