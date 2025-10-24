package mappers

import (
	"fleetmanagement/api/controller/pb"
	"fleetmanagement/logic/model"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertProtobufVinToModelVin(vin *pb.Vin) model.Vin {
	return model.Vin{
		Vin: vin.Vin,
	}
}

func ConvertProtobufTimeStampToDate(timestamp *timestamppb.Timestamp) time.Time {
	time, _ := time.Parse(time.DateOnly, timestamp.AsTime().Format("2006-01-02"))
	return time
}
