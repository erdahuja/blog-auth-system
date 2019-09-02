package models

import "github.com/jinzhu/gorm"

// User model represented in db
type User struct {
	gorm.Model
	Email             string `gorm:"unique_index;not null"`
	Password          string `gorm:"-"`
	PasswordHash      string `gorm:"not null"`
	RememberToken     string `gorm:"-"`
	RememberTokenHash string `gorm:"unique_index;not null"`
}
