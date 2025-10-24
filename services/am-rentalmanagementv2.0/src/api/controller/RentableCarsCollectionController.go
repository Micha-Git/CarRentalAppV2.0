package controller

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"rentalmanagement/api/controller/mappers"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/operations"
)

type RentableCarsCollectionController struct {
	ops operations.RentableCarsCollectionOperationsInterface
	pb.UnimplementedRentableCarsCollectionServiceServer
}

func NewRentableCarsCollectionController(ops operations.RentableCarsCollectionOperationsInterface) RentableCarsCollectionController {
	return RentableCarsCollectionController{ops: ops}
}

// Implement the AddCarToRental RPC method
func (controller RentableCarsCollectionController) AddCarToRental(ctx context.Context, req *pb.AddCarToRentalRequest) (*pb.AddCarToRentalResponse, error) {
	log.Info("Starting to add a car to a rental.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.AddCarToRentalResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			},
		}, nil
	}

	if req.Vin == nil || len(req.Location) == 0 {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Vin, or location : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.AddCarToRentalResponse{
			Error: errorDetail,
		}, nil
	}

	pbRentableCar, err := controller.ops.AddCarToRental(mappers.ConvertProtobufVinToModelVin(req.Vin), req.Location)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.AddCarToRentalResponse{
			Error: errorDetail,
		}, nil
	}

	rentableCar := mappers.ConvertModelRentableCarToProtobufRentableCar(pbRentableCar)

	return &pb.AddCarToRentalResponse{
		Car: rentableCar,
	}, nil
}

func (controller RentableCarsCollectionController) RemoveRentableCar(ctx context.Context, req *pb.RemoveRentableCarRequest) (*pb.RemoveRentableCarResponse, error) {
	log.Info("Starting to remove a rentable car.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.RemoveRentableCarResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			},
		}, nil
	}

	if req.Vin == nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Vin : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.RemoveRentableCarResponse{
			Error: errorDetail,
		}, nil
	}

	err := controller.ops.RemoveRentableCar(mappers.ConvertProtobufVinToModelVin(req.Vin))
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.RemoveRentableCarResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.RemoveRentableCarResponse{
		Result: true,
	}, nil
}
