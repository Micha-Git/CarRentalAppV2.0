package entities

type RentableCarPersistenceEntity struct {
	Vin         string  `gorm:"type:text;primary_key;"`
	Brand       string  `gorm:"type:text;"`
	Model       string  `gorm:"type:text;"`
	Location    string  `gorm:"type:text;"`
	PricePerDay float32 `gorm:"type:float;"`
}

// TableName sets the insert table name for this struct type
func (p *RentableCarPersistenceEntity) TableName() string {
	return "rentable_car"
}
