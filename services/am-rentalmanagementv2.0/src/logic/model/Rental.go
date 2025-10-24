package model

import (
	"time"
)

// The model of the car that is used by all internal operations
// This does not correspond to the API Diagram
type Rental struct {
	Id         string
	StartDate  time.Time
	EndDate    time.Time
	Car        RentableCar
	Price      float32
	CustomerId string
}
