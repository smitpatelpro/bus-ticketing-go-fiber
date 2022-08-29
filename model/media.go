package model

import "gorm.io/gorm"

// Model struct
type Media struct {
	gorm.Model
	Path string `gorm:"not null" json:"path"`
}
