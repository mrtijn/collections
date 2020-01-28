package main

import (
	"collections/database"
	"collections/models"
	"collections/routes"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// var tmdbConfig = tmdb.GetTmdbConfiguration()

func main() {
	// Port
	port := 3000

	// Setup db
	var database = database.Connect()
	database.AutoMigrate(&models.User{}, &models.Collection{})
	database.Model(&models.Collection{}).AddForeignKey("author_id", "users(id)", "RESTRICT", "CASCADE")

	router := gin.Default()

	// Setup sessions
	store := cookie.NewStore([]byte(os.Getenv("sessionSecret")))
	router.Use(sessions.Sessions("collections", store))

	// Setup TMDB Connection

	// Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("x-custom-header")
	router.Use(cors.New(corsConfig))

	// Setup routes
	r := router.Group("/v1")
	routes.CreateRoutes(r)

	// Start server
	startWebServer(router, port)
}

func startWebServer(server *gin.Engine, port int) {
	fmt.Println("Started webserver on port", port)
	server.Run(":" + strconv.Itoa(port))

}
