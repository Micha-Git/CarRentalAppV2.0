package operations

import (
	"rentalmanagement/logic/model"
)

type CustomerOperations struct {
	rentalRepository      model.RentalRepositoryInterface
	rentableCarRepository model.RentableCarRepositoryInterface
}

func NewCustomerOperations(rentalRepository model.RentalRepositoryInterface, rentableCarRepository model.RentableCarRepositoryInterface) CustomerOperations {
	return CustomerOperations{rentalRepository: rentalRepository, rentableCarRepository: rentableCarRepository}
}
