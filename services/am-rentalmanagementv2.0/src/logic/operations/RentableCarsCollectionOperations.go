package operations

import (
	"rentalmanagement/logic/model"
)

type RentableCarsCollectionOperations struct {
	rentalRepository      model.RentalRepositoryInterface
	rentableCarRepository model.RentableCarRepositoryInterface
	carRepository         model.CarRepositoryInterface
}

func NewRentableCarsCollectionOperations(rentalRepository model.RentalRepositoryInterface, rentableCarRepository model.RentableCarRepositoryInterface, carRepository model.CarRepositoryInterface) RentableCarsCollectionOperations {
	return RentableCarsCollectionOperations{rentalRepository: rentalRepository, rentableCarRepository: rentableCarRepository, carRepository: carRepository}
}
