package controller

import (
	"context"
	"fleetmanagement/api/controller/mappers"
	"fleetmanagement/api/controller/pb"
	"fleetmanagement/logic/model"
	"fleetmanagement/logic/operations"
	"fmt"

	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
)

type FleetController struct {
	ops operations.FleetOperationsInterface
	pb.UnimplementedFleetServiceServer
}

func NewFleetController(ops operations.FleetOperationsInterface) FleetController {
	return FleetController{ops: ops}
}

// Implement the AddCarToFleet RPC method
func (controller FleetController) AddCarToFleet(ctx context.Context, req *pb.AddCarToFleetRequest) (*pb.AddCarToFleetResponse, error) {
	log.Info("Starting to add a new car to the fleet.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.AddCarToFleetResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.Vin.GetVin() == "" || req.GetFleetId() == "" || req.GetLocation() == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("VIN, fleetId or location is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.AddCarToFleetResponse{
			Error: errorDetail,
		}, nil
	}

	vin := model.Vin{
		Vin: req.Vin.GetVin(),
	}

	car, err := controller.ops.AddCarToFleet(vin, req.GetFleetId(), req.GetLocation())
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.AddCarToFleetResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.AddCarToFleetResponse{
		Car: mappers.ConvertModelCarToProtobufCar(car),
	}, nil
}

// Implement the ListCarsInFleet RPC method
func (controller FleetController) ListCarsInFleet(ctx context.Context, req *pb.ListCarsInFleetRequest) (*pb.ListCarsInFleetResponse, error) {
	log.Info("Starting to list cars in a fleet.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListCarsInFleetResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.GetFleetId() == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("FleetId is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListCarsInFleetResponse{
			Error: errorDetail,
		}, nil
	}
	cars, err := controller.ops.ListCarsInFleet(req.GetFleetId())
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListCarsInFleetResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.ListCarsInFleetResponse{
		Cars: mappers.ConvertModelCarsToProtobufCars(cars),
	}, nil
}

// Implement the RemoveCarFromFleet RPC method
func (controller FleetController) RemoveCarFromFleet(ctx context.Context, req *pb.RemoveCarFromFleetRequest) (*pb.RemoveCarFromFleetResponse, error) {
	log.Info("Starting to remove a car from the fleet.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.RemoveCarFromFleetResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.Vin.GetVin() == "" {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("VIN is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.RemoveCarFromFleetResponse{
			Error: errorDetail,
		}, nil
	}

	vin := model.Vin{
		Vin: req.Vin.GetVin(),
	}

	// Call the relevant logic operation to remove the car.
	removal, err := controller.ops.RemoveCarFromFleet(vin)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.RemoveCarFromFleetResponse{
			Error:  errorDetail,
			Result: false,
		}, nil
	}

	// If successful, return a response confirming the cancellation.
	return &pb.RemoveCarFromFleetResponse{Result: removal}, nil
}
