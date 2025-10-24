package external

import (
	"fleetmanagement/logic/model"
)

type RentalManagementAPIStub struct {
	cars []model.Car
}

func NewRentalManagementAPIStub(cars []model.Car) *RentalManagementAPIStub {
	return &RentalManagementAPIStub{cars: cars}
}

func (r *RentalManagementAPIStub) AddCarToRental(vin model.Vin, location string) (model.RentableCar, error) {

	for _, car := range r.cars {
		if car.Vin == vin {
			resp := model.RentableCar{
				Vin:         car.Vin,
				Brand:       car.Brand,
				Model:       car.Model,
				Location:    location,
				PricePerDay: 50,
			}

			return resp, nil
		}
	}

	return model.RentableCar{}, nil
}

func (r *RentalManagementAPIStub) RemoveRentableCar(vin model.Vin) (bool, error) {

	for _, car := range r.cars {
		if car.Vin == vin {

			return true, nil
		}
	}

	return false, nil

}
