package entity

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
	Description string `json:"description"`
	CategoryID uint      `json:"category_id"`
	Category   Category  `json:"category" gorm:"foreignKey:CategoryID"`
	Price       float64   `json:"price" gorm:"not null"`
	Capacity    int       `json:"capacity" gorm:"not null"`
	Status      string    `json:"status" gorm:"type:enum('active','ongoing','finished');default:'active'"`
	StartTime   time.Time `json:"start_time"`
}