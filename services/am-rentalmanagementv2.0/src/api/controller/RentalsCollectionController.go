package controller

import (
	"context"
	"fmt"
	"rentalmanagement/api/controller/mappers"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/operations"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc/codes"
)

type RentalsCollectionController struct {
	ops operations.RentalsCollectionOperationsInterface
	pb.UnimplementedRentalsCollectionServiceServer
}

func NewRentalsCollectionController(ops operations.RentalsCollectionOperationsInterface) RentalsCollectionController {
	return RentalsCollectionController{ops: ops}
}

// Implement the ListAvailableCars RPC method
func (controller RentalsCollectionController) ListAvailableCars(ctx context.Context, req *pb.ListAvailableCarsRequest) (*pb.ListAvailableCarsResponse, error) {
	log.Info("Starting to list available cars.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListAvailableCarsResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.StartDate == nil || req.EndDate == nil || len(req.Location) == 0 {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Start date, end date, or location is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListAvailableCarsResponse{
			Error: errorDetail,
		}, nil
	}

	availableCars, err := controller.ops.ListAvailableCars(mappers.ConvertProtobufTimeStampToDate(req.StartDate), mappers.ConvertProtobufTimeStampToDate(req.EndDate), req.Location)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListAvailableCarsResponse{
			Error: errorDetail,
		}, nil
	}

	var pbAvailableCars []*pb.RentableCar
	for _, availableCar := range availableCars {
		pbAvailableCars = append(pbAvailableCars, mappers.ConvertModelRentableCarToProtobufRentableCar(availableCar))
	}

	return &pb.ListAvailableCarsResponse{
		Cars: pbAvailableCars,
	}, nil
}

func (controller RentalsCollectionController) ListCarRentals(ctx context.Context, req *pb.ListCarRentalsRequest) (*pb.ListCarRentalsResponse, error) {
	log.Info("Starting to list car rentals.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListCarRentalsResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if req.Vin == nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Vin is not valid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListCarRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	rentals, err := controller.ops.ListCarRentals(mappers.ConvertProtobufVinToModelVin(req.Vin))
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListCarRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	var pbRentals []*pb.Rental
	for _, rental := range rentals {
		pbRentals = append(pbRentals, mappers.ConvertModelRentalToProtobufRental(rental))
	}

	return &pb.ListCarRentalsResponse{
		Rentals: pbRentals,
	}, nil
}

func (controller RentalsCollectionController) ListCustomerRentals(ctx context.Context, req *pb.ListCustomerRentalsRequest) (*pb.ListCustomerRentalsResponse, error) {
	log.Info("Starting to list customer rentals.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ListCustomerRentalsResponse{
			Error: &pb.ErrorDetail{
				Message: codes.Unknown.String(),
			}}, nil
	}

	// Validate the request data. If validation fails, return an error.
	if len(req.CustomerId) == 0 {
		errorDetail := &pb.ErrorDetail{
			Message: codes.InvalidArgument.String(),
			Details: fmt.Sprintf("Customer ID is invalid : %s", req),
		}
		log.Errorf("Error: %s - %s", req, errorDetail.Details)
		return &pb.ListCustomerRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	rentals, err := controller.ops.ListCustomerRentals(req.CustomerId)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ListCustomerRentalsResponse{
			Error: errorDetail,
		}, nil
	}

	var pbRentals []*pb.Rental
	for _, rental := range rentals {
		pbRentals = append(pbRentals, mappers.ConvertModelRentalToProtobufRental(rental))
	}

	return &pb.ListCustomerRentalsResponse{
		Rentals: pbRentals,
	}, nil
}
