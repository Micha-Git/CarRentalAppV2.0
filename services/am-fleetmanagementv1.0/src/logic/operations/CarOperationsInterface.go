package operations

import (
	"fleetmanagement/logic/model"
)

type CarOperationsInterface interface {
	ViewCarInformation(vin model.Vin) (model.Car, error)
}
