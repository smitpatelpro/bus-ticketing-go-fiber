package model

import "gorm.io/gorm"

// Model struct
type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Amount      int    `gorm:"not null" json:"amount"`
}
