package operations

import (
	"fleetmanagement/infrastructure/external"
	"fleetmanagement/logic/model"
)

type FleetOperations struct {
	repository          model.PostgresRepositoryInterface
	carApi              external.CarAPIInterface
	rentalManagementApi external.RentalManagementAPIInterface
}

func NewFleetOperations(repository model.PostgresRepositoryInterface, carApi external.CarAPIInterface, rentalManagementApi external.RentalManagementAPIInterface) FleetOperations {
	return FleetOperations{repository: repository, carApi: carApi, rentalManagementApi: rentalManagementApi}
}
