package entities

type CarPersistenceEntity struct {
	Vin      string `gorm:"type:text;primary_key;"`
	Model    string `gorm:"type:text;"`
	Brand    string `gorm:"type:text;"`
	FleetId  string `gorm:"type:text;"`
	Location string `gorm:"type:text;"`
}

// TableName sets the insert table name for this struct type
func (p *CarPersistenceEntity) TableName() string {
	return "car"
}
