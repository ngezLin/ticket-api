package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"not null;unique"`
	Password string `json:"password" gorm:"not null"`
	Role string `json:"role" gorm:"type:ENUM('admin', 'customer');not null"`
	Balance  float64 `json:"balance"`
}