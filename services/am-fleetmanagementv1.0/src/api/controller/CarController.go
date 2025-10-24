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

type CarController struct {
	ops operations.CarOperationsInterface
	pb.UnimplementedCarServiceServer
}

func NewCarController(ops operations.CarOperationsInterface) CarController {
	return CarController{ops: ops}
}

// Implement the ViewCarInformation RPC method
func (controller CarController) ViewCarInformation(ctx context.Context, req *pb.ViewCarInformationRequest) (*pb.ViewCarInformationResponse, error) {
	log.Info("Starting to view car information.")

	if req == nil {
		log.Errorf("Error: empty request")
		return &pb.ViewCarInformationResponse{
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
		return &pb.ViewCarInformationResponse{
			Error: errorDetail,
		}, nil
	}

	vin := model.Vin{
		Vin: req.Vin.GetVin(),
	}

	car, err := controller.ops.ViewCarInformation(vin)
	if err != nil {
		errorDetail := &pb.ErrorDetail{
			Message: codes.Internal.String(),
			Details: err.Error(),
		}
		return &pb.ViewCarInformationResponse{
			Error: errorDetail,
		}, nil
	}

	return &pb.ViewCarInformationResponse{
		Car: mappers.ConvertModelCarToProtobufCar(car),
	}, nil
}
