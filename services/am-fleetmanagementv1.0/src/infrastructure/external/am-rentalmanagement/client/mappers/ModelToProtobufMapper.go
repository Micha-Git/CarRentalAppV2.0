package mappers

import (
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client/pb"
	"fleetmanagement/logic/model"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertModelRentableCarToProtobufRentableCar(rentableCar model.RentableCar) *pb.RentableCar {
	return &pb.RentableCar{
		Vin:         ConvertModelVinToProtobufVin(rentableCar.Vin),
		Brand:       rentableCar.Brand,
		Model:       rentableCar.Model,
		Location:    rentableCar.Location,
		PricePerDay: rentableCar.PricePerDay,
	}
}

/* Not required by fleet management

func ConvertModelRentalToProtobufRental(rental model.Rental) *pb.Rental {
	return &pb.Rental{
		Id:         rental.Id,
		StartDate:  ConvertModelDateToProtobufTimestamp(rental.StartDate),
		EndDate:    ConvertModelDateToProtobufTimestamp(rental.EndDate),
		Car:        ConvertModelRentableCarToProtobufRentableCar(rental.Car),
		Price:      rental.Price,
		CustomerId: rental.CustomerId,
	}
}
*/

func ConvertModelVinToProtobufVin(vin model.Vin) *pb.Vin {
	return &pb.Vin{
		Vin: vin.Vin,
	}
}

func ConvertModelDateToProtobufTimestamp(date time.Time) *timestamppb.Timestamp {
	return timestamppb.New(date)
}
