package models

import "time"

type Subscription struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"not null;index;constraint:OnDelete:CASCADE"`
	StartDate time.Time `gorm:"not null"`
	EndDate   time.Time `gorm:"not null;index"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
}

func (Subscription) TableName() string {
	return "subscriptions"
}
