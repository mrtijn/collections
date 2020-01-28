package collection

import (
	"collections/database"
	"collections/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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

	collectionForUser := db.Where("author_id = ?", 55).First(collection)
	if collectionForUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "You already have a collection with the name " + collection.Name})
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

	c.JSON(http.StatusOK, collection)
}

// GetMyCollections Get collections for loggedin user
func GetMyCollections(c *gin.Context) {
	collection := &models.Collection{}

	session := sessions.Default(c)
	id := session.Get("userID")

	err := db.Where("author_id = ?", id).Find(collection).Error

	if err != nil {
		c.JSON(http.StatusNotFound, "Collection not found")
		return
	}

	c.JSON(http.StatusOK, collection)
}
