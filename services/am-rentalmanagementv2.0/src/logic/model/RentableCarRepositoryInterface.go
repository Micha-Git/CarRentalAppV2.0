package model

type RentableCarRepositoryInterface interface {
	AddRentableCar(rentableCar RentableCar) (RentableCar, error)
	GetRentableCar(vin Vin) (RentableCar, error)
	ListRentableCarsByLocation(location string) ([]RentableCar, error)
	RemoveRentableCar(vin string) error
}
