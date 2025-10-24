package model

// The model of the fleet that is used by all internal operations
// This corresponds to the API Diagram
type Fleet struct {
	Cars         []Car
	Location     string
	FleetManager string
	FleetId      string
}
