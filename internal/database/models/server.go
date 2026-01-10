package models

import "time"

type Server struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"uniqueIndex;not null"`
	Host        string    `gorm:"not null"`
	IsActive    bool      `gorm:"default:true;index"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`

	Inbounds []Inbound `gorm:"foreignKey:ServerID"`
}

func (Server) TableName() string {
	return "servers"
}
