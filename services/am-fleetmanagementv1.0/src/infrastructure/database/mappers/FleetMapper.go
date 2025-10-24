package mappers

import (
	"fleetmanagement/infrastructure/database/entities"
	"fleetmanagement/logic/model"
)

func ConvertFleetToFleetPersistenceEntity(fleet model.Fleet) entities.FleetPersistenceEntity {

	return entities.FleetPersistenceEntity{
		FleetId:      fleet.FleetId,
		Location:     fleet.Location,
		FleetManager: fleet.FleetManager,
	}
}

func ConvertFleetPersistenceEntityToFleet(fleetPers entities.FleetPersistenceEntity) model.Fleet {

	return model.Fleet{
		FleetId:      fleetPers.FleetId,
		Location:     fleetPers.Location,
		FleetManager: fleetPers.FleetManager,
	}
}
