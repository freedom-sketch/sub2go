package models

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UserID    int64     `gorm:"column:user_id;unique;not null;index"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	Subscriptions []Subscription `gorm:"foreignKey:UserID;references:UserID;constraint:OnDelete:CASCADE"`
	Admin         *Admin         `gorm:"foreignKey:UserID;references:UserID"`
}

func (User) TableName() string {
	return "users"
}
