package models

import "time"

type Inbound struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	ServerID    uint   `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	Tag         string `gorm:"uniqueIndex;not null"`
	Protocol    string `gorm:"not null"`
	Port        int    `gorm:"not null;index"`
	Network     string `gorm:"not null"`
	Security    string `gorm:"default:reality"`
	Flow        string `gorm:"default:''"`
	ShortIds    string `gorm:"type:text"`
	PublicKey   string
	PrivateKey  string
	Target      string
	SNI         string
	IsActive    bool      `gorm:"default:true;index"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
	Description string    `gorm:"type:text"`

	Server Server `gorm:"-"`
}

func (Inbound) TableName() string {
	return "inbounds"
}
