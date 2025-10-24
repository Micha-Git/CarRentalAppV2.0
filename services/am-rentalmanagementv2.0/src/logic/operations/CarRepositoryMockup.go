package operations

import (
	"fmt"
	"net/http"
	"rentalmanagement/logic/model"
)

type CarRepositoryMockup struct {
	cars []model.Car
}

func NewCarRepositoryMockup(cars []model.Car) model.CarRepositoryInterface {
	return &CarRepositoryMockup{
		cars: cars,
	}
}

func (carRepo *CarRepositoryMockup) GetCar(vin model.Vin) (model.Car, error) {
	for _, car := range carRepo.cars {
		if car.Vin == vin {
			return car, nil
		}
	}
	return model.Car{}, fmt.Errorf("API request failed with status code: %d", http.StatusBadRequest)
}
