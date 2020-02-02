package collection

import (
	"collections/database"
	"collections/models"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var db = database.Connect()

// CreateCollection creates a collection
func CreateCollection(c *gin.Context) {
	collection := &models.Collection{}

	err := c.Bind(collection)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data"})
		return
	}
	// Get user id from session
	session := sessions.Default(c)
	id := session.Get("userID")

	collection.AuthorID = id.(uint)

	err = db.Where("author_id = ? AND title = ?", id, collection.Title).First(collection).Error

	if !gorm.IsRecordNotFoundError(err) {
		c.JSON(http.StatusConflict, gin.H{"error": "You already have a collection with the title " + collection.Title})
		return
	}

	err = db.Create(collection).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Could not create collection"})
		return
	}

	c.JSON(http.StatusOK, collection)
}

// GetCollection gets collection data
func GetCollection(c *gin.Context) {
	collection := &models.Collection{}

	var id = c.Param("id")

	err := db.Where("Id = ?", id).First(collection).Error

	if err != nil {
		c.JSON(http.StatusNotFound, "Collection not found")
		return
	}

	db.Preload("Movies").Find(&collection)

	c.JSON(http.StatusOK, collection)
}

// GetMyCollections Get collections for loggedin user
func GetMyCollections(c *gin.Context) {
	collections := []models.Collection{}

	session := sessions.Default(c)
	id := session.Get("userID")

	err := db.Where("author_id = ?", id).Find(&collections).Error

	if err != nil {
		c.JSON(http.StatusNotFound, "Collection not found")
		return
	}

	c.JSON(http.StatusOK, collections)
}

// AddToCollectionPayload model
type AddToCollectionPayload struct {
	CollectionID string `json:"collection_id"`
	MovieID      string `json:"movie_id"`
}

// AddToCollection add movie to collection
func AddToCollection(c *gin.Context) {
	payload := AddToCollectionPayload{}

	err := c.Bind(&payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	collection := models.Collection{}
	movie := models.Movie{}
	err = db.Where("id = ?", payload.CollectionID).First(&collection).Error
	if err != nil {
		c.JSON(http.StatusNotFound, "Collection not found")
	}

	err = db.Where("id = ?", payload.MovieID).First(&movie).Error
	if err != nil {
		c.JSON(http.StatusNotFound, "Movie not found")
	}

	db.Model(&collection).Association("Movies").Append([]models.Movie{movie})

	c.JSON(http.StatusOK, "Collection created")
}
