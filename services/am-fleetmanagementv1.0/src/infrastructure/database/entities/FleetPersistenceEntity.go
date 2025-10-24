package entities

type FleetPersistenceEntity struct {
	FleetId      string `gorm:"type:text;primary_key;"`
	Location     string `gorm:"type:text;primary_key;"`
	FleetManager string `gorm:"type:text;"`
}

// TableName sets the insert table name for this struct type
func (p *FleetPersistenceEntity) TableName() string {
	return "fleet"
}
