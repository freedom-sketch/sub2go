package models

type Server struct {
	ID          uint   `gorm:"primaryKey;autoIncrement"`
	IPAddress   string `gorm:"column:ip_address;type:text;not null"`
	Port        int    `gorm:"default:4443"`
	PanelURL    string `gorm:"column:panel_url;type:text"`
	InboundID   int    `gorm:"default:1"`
	Login       string `gorm:"type:text"`
	Password    string `gorm:"type:text"`
	IsActive    bool   `gorm:"default:true;index"`
	CountryID   uint   `gorm:"not null;constraint:OnDelete:CASCADE"`
	Description string `gorm:"type:text"`

	Country Country `gorm:"foreignKey:CountryID"`
}

func (Server) TableName() string {
	return "servers"
}
