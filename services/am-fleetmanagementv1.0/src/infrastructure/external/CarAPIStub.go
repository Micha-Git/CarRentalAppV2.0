package external

import (
	"fleetmanagement/logic/model"
)

type CarAPIStub struct {
	cars []model.Car
}

func NewCarAPIStub(cars []model.Car) *CarAPIStub {
	return &CarAPIStub{cars: cars}
}

func (c *CarAPIStub) GetCar(vin model.Vin) (model.Car, error) {

	for _, car := range c.cars {
		if car.Vin == vin {
			return car, nil
		}
	}

	return model.Car{}, nil
}
