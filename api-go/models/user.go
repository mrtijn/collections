package models

import "github.com/jinzhu/gorm"

// User model
type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"type:varchar(100);unique;"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
