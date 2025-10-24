package mappers

import (
	"fleetmanagement/api/controller/pb"
	"fleetmanagement/logic/model"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertModelVinToProtobufVin(vin model.Vin) *pb.Vin {
	return &pb.Vin{
		Vin: vin.Vin,
	}
}

func ConvertModelCarToProtobufCar(car model.Car) *pb.Car {
	return &pb.Car{
		Vin:   &pb.Vin{Vin: car.Vin.Vin},
		Model: car.Model,
		Brand: car.Brand}
}

func ConvertModelCarsToProtobufCars(cars []model.Car) []*pb.Car {
	var protobufCars []*pb.Car
	for _, car := range cars {
		protobufCars = append(protobufCars, ConvertModelCarToProtobufCar(car))
	}
	return protobufCars
}

func ConvertModelFleetToProtobufFleet(fleet model.Fleet) *pb.Fleet {
	return &pb.Fleet{
		FleetId:      fleet.FleetId,
		Cars:         ConvertModelCarsToProtobufCars(fleet.Cars),
		Location:     fleet.Location,
		FleetManager: fleet.FleetManager}
}

func ConvertModelDateToProtobufTimestamp(date time.Time) *timestamppb.Timestamp {
	return timestamppb.New(date)
}
