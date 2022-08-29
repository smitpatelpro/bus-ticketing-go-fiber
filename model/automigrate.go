package model

import (
	"fmt"

	"gorm.io/gorm"
)

func RunAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	fmt.Println("User Database Migrated")
	db.AutoMigrate(&BusOperatorProfile{})
	fmt.Println("BusOperatorProfile Database Migrated")
	db.AutoMigrate(&Media{})
	fmt.Println("Media Database Migrated")
}
