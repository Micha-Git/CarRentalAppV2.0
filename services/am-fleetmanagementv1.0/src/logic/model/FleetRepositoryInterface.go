package model

type FleetRepositoryInterface interface {
	GetFleet(fleetId string) (Fleet, error)
}
