package handler

import (
	"bus-api/model"
)

// Model struct
type User struct {
	Username    string     `gorm:"unique_index;not null" json:"username"`
	Email       string     `gorm:"unique_index;not null" json:"email"`
	Password    string     `gorm:"not null" json:"password"`
	FullName    string     `json:"full_name"`
	PhoneNumber string     `json:"phone_number"`
	Role        model.Role `json:"role"`
	IsAdmin     bool       `json:"is_admin"`
}
