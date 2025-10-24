package mappers

import (
	"fleetmanagement/infrastructure/external/am-rentalmanagement/client/pb"
	"fleetmanagement/logic/model"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertProtobufRentableCarToModelRentableCar(rentableCar *pb.RentableCar) model.RentableCar {
	return model.RentableCar{
		Vin:         ConvertProtobufVinToModelVin(rentableCar.Vin),
		Brand:       rentableCar.Brand,
		Model:       rentableCar.Model,
		Location:    rentableCar.Location,
		PricePerDay: rentableCar.PricePerDay,
	}
}

/* Not required by fleet management

func ConvertProtobufRentalToModelRental(rental *pb.Rental) model.Rental {
	return model.Rental{
		Id:         rental.Id,
		StartDate:  ConvertProtobufTimeStampToDate(rental.StartDate),
		EndDate:    ConvertProtobufTimeStampToDate(rental.EndDate),
		Car:        ConvertProtobufRentableCarToModelRentableCar(rental.Car),
		Price:      rental.Price,
		CustomerId: rental.CustomerId,
	}
}
*/

func ConvertProtobufVinToModelVin(vin *pb.Vin) model.Vin {
	return model.Vin{
		Vin: vin.Vin,
	}
}

func ConvertProtobufTimeStampToDate(timestamp *timestamppb.Timestamp) time.Time {
	time, _ := time.Parse(time.DateOnly, timestamp.AsTime().Format("2006-01-02"))
	return time
}
