package models

import "github.com/jinzhu/gorm"

// Movie model
type Movie struct {
	gorm.Model
	Title       string       `json:"title" gorm:"type:varchar(100);not null;"`
	TmdbID      string       `json:"tmdb_id" gorm:"not null;unique;"`
	Details     MovieDetails `json:"details"`
	Collections []Collection `json:"collections" gorm:"many2many:movie_collections"`
}

// MovieDetails model
type MovieDetails struct {
	Title        string `json:"title"`
	Overview     string `json:"overview"`
	ReleaseDate  string `json:"release_date"`
	PosterPath   string `json:"poster_path"`
	BackdropPath string `json:"backdrop_path"`
}
