package model

//.
type CarRepositoryInterface interface {
	AddCarToFleet(car Car, fleetId string, location string) (Car, error)
	GetCar(vin Vin) (Car, error)
	RemoveCar(vin Vin) (bool, error)
	ListAllCars(fleetId string) ([]Car, error)
}
