package operations

import (
	"rentalmanagement/logic/model"
)

type RentalsCollectionOperations struct {
	rentalRepository      model.RentalRepositoryInterface
	rentableCarRepository model.RentableCarRepositoryInterface
}

func NewRentalsCollectionOperations(rentalRepository model.RentalRepositoryInterface, rentableCarRepository model.RentableCarRepositoryInterface) RentalsCollectionOperations {
	return RentalsCollectionOperations{rentalRepository: rentalRepository, rentableCarRepository: rentableCarRepository}
}
