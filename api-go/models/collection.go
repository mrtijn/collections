package models

import "github.com/jinzhu/gorm"

// Collection model
type Collection struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(100);not null;"`
	Type     string `json:"type" gorm:"not null"`
	Private  bool   `json:"private" gorm:"not null"`
	AuthorID uint   `json:"author_id"`
	Author   User   `gorm:"foreignKey:author_id"`
}
