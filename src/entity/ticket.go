package entity

import "gorm.io/gorm"

type Ticket struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	EventID  uint   `json:"event_id"`
	Quantity int    `json:"quantity"`
	TotalPrice float64 `json:"total_price"` 
	Status   string `json:"status" gorm:"type:enum('active','cancelled')"`

	User  User  `json:"-" gorm:"foreignKey:UserID"`
	Event Event `json:"event" gorm:"foreignKey:EventID"`
}