package mappers

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"rentalmanagement/api/controller/pb"
	"rentalmanagement/logic/model"
	"time"
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

func ConvertProtobufVinToModelVin(vin *pb.Vin) model.Vin {
	return model.Vin{
		Vin: vin.Vin,
	}
}

func ConvertProtobufTimeStampToDate(timestamp *timestamppb.Timestamp) time.Time {
	time, _ := time.Parse(time.DateOnly, timestamp.AsTime().Format("2006-01-02"))
	return time
}
