package model

import "gorm.io/gorm"

type Role string

const (
	ROLE_ADMIN        Role = "ADMIN"
	ROLE_BUS_OPERATOR Role = "BUS_OPERATOR"
	ROLE_CUSTOMER     Role = "CUSTOMER"
)

// Model struct
type User struct {
	gorm.Model
	Username    string `gorm:"unique_index;not null" json:"username"`
	Email       string `gorm:"unique_index;not null" json:"email"`
	Password    string `gorm:"not null" json:"-"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Role        Role   `json:"role"`
	IsAdmin     bool   `json:"is_admin"`
}
