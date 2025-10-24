package external

import (
	"fleetmanagement/logic/model"
)

// .
type CarAPIInterface interface {
	GetCar(vin model.Vin) (model.Car, error)
}
