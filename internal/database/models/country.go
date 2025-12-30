package models

type Country struct {
	ID   uint   `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"type:text;not null"`
	Code string `gorm:"type:text;unique;not null"`

	Servers []Server `gorm:"foreignKey:CountryID"`
}

func (Country) TableName() string {
	return "countries"
}
