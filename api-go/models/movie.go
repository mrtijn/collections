package models

import "github.com/jinzhu/gorm"

// Movie model
type Movie struct {
	gorm.Model
	Title  string `json:"name" gorm:"type:varchar(100);not null;"`
	TmdbID string
}
