package movie

import (
	"collections/database"
	"collections/models"
	"collections/pkg/tmdb"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var db = database.Connect()

// CreateMovie create movie
func CreateMovie(c *gin.Context) {
	movie := &models.Movie{}

	err := c.Bind(movie)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data"})
		return
	}

	err = db.Create(movie).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, movie)
}

// SearchMovie search movie
func SearchMovie(c *gin.Context) {
	searchTerm := c.Param("searchterm")

	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, "Provide searchterm")
	}
	resp, err := tmdb.SearchMovie(searchTerm)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, "Something didnt return")
		return
	}

	c.JSON(http.StatusOK, resp)
}

// func GetMovieById(c.) {

// }
