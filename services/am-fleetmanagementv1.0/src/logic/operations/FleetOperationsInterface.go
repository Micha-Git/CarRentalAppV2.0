package operations

import (
	"fleetmanagement/logic/model"
)

type FleetOperationsInterface interface {
	AddCarToFleet(vin model.Vin, fleetId string, location string) (model.Car, error)
	ListCarsInFleet(fleetId string) ([]model.Car, error)
	RemoveCarFromFleet(vin model.Vin) (bool, error)
}
