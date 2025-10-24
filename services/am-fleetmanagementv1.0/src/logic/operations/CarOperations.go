package operations

import (
	"fleetmanagement/infrastructure/external"
	"fleetmanagement/logic/model"
)

type CarOperations struct {
	repository model.PostgresRepositoryInterface
	carApi     external.CarAPIInterface
}

func NewCarOperations(repository model.PostgresRepositoryInterface, carApi external.CarAPIInterface) CarOperations {
	return CarOperations{repository: repository, carApi: carApi}
}
