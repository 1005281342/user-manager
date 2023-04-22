package entity

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}
