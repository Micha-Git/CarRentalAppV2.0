package external

import (
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client"
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client/pb"
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type RentalManagementAPI struct {
	client *client.RentableCarsCollectionClient
}

func NewRentalManagementAPI(rpcClient pb.RentableCarsCollectionServiceClient) *RentalManagementAPI {
	return &RentalManagementAPI{
		client: client.NewRentableCarsCollectionClient(rpcClient),
	}
}

func (r *RentalManagementAPI) AddCarToRental(vin model.Vin, location string) (model.RentableCar, error) {
	var msg string

	resp, err := r.client.AddCarToRental(vin, location)
	if err != nil {
		msg = "Failed to make RPC AddCarToRental request to the Rental Management API"
		log.Error(msg, ": ", err)
		return model.RentableCar{}, fmt.Errorf("%s: %w", msg, err)
	}

	return resp, nil
}

func (r *RentalManagementAPI) RemoveRentableCar(vin model.Vin) (bool, error) {
	var msg string

	resp, err := r.client.RemoveRentableCar(vin)
	if err != nil {
		msg = "Failed to make RPC RemoveRentableCar request to the Rental Management API"
		log.Error(msg, ": ", err)
		return false, fmt.Errorf("%s: %w", msg, err)
	}

	return resp, nil
}
