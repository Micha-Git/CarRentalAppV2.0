package model

// The model of the RenalCar that is used by all internal operations
// This corresponds to the API Diagram
type RentableCar struct {
	Vin         Vin
	Brand       string
	Model       string
	Location    string
	PricePerDay float32
}
