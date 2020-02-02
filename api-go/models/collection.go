package models

import "github.com/jinzhu/gorm"

// Collection model
type Collection struct {
	gorm.Model
	Title    string  `json:"title" gorm:"type:varchar(100);not null;"`
	Type     string  `json:"type" gorm:"not null"`
	Private  bool    `json:"private" gorm:"not null"`
	AuthorID uint    `json:"author_id"`
	Movies   []Movie `gorm:"many2many:movie_collections"`
}
