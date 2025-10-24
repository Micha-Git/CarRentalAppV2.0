package client

import (
	"context"
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client/mappers"
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client/pb"
	"fleetmanagement/logic/model"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type RentableCarsCollectionClient struct {
	client pb.RentableCarsCollectionServiceClient
}

func NewRentableCarsCollectionClient(client pb.RentableCarsCollectionServiceClient) *RentableCarsCollectionClient {
	return &RentableCarsCollectionClient{
		client: client,
	}
}

// Sends an RPC request to add a car to a rental
func (c *RentableCarsCollectionClient) AddCarToRental(vin model.Vin, location string) (model.RentableCar, error) {
	log.Info("Sending request to add car to rental.")

	protobufVin := mappers.ConvertModelVinToProtobufVin(vin)

	req := &pb.AddCarToRentalRequest{
		Vin:      protobufVin,
		Location: location,
	}

	resp, err := c.client.AddCarToRental(context.Background(), req)
	if err != nil {
		log.Errorf("Error while calling AddCarToRental: %v", err)
		return model.RentableCar{}, fmt.Errorf("error while calling AddCarToRental: %w", err)
	}

	rentableCar := mappers.ConvertProtobufRentableCarToModelRentableCar(resp.Car)

	return rentableCar, nil
}

// Sends an RPC request to remove a rentable car
func (c *RentableCarsCollectionClient) RemoveRentableCar(vin model.Vin) (bool, error) {
	log.Info("Sending request to remove a rentable car.")

	protobufVin := mappers.ConvertModelVinToProtobufVin(vin)

	req := &pb.RemoveRentableCarRequest{
		Vin: protobufVin,
	}

	resp, err := c.client.RemoveRentableCar(context.Background(), req)
	if err != nil {
		log.Errorf("Error while calling RemoveRentableCar: %v", err)
		return false, fmt.Errorf("error while calling RemoveRentableCar: %w", err)
	}

	return resp.Result, nil
}
