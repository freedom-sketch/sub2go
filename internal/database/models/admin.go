package models

type Admin struct {
	ID     uint   `gorm:"primaryKey;autoIncrement"`
	Name   string `gorm:"type:text;unique"`
	UserID int64  `gorm:"not null;index"`

	User User `gorm:"foreignKey:UserID;references:UserID"`
}

func (Admin) TableName() string {
	return "admins"
}
