package main

import (
	"collections/database"
	"collections/models"
	"collections/routes"
	"fmt"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Port
	port := 3000

	// Setup db
	var database = database.Connect()
	database.AutoMigrate(&models.User{})

	router := gin.Default()

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
	// fmt.Println("Starting webserver...")
	fmt.Println("Started webserver on port", port)
	server.Run(":" + strconv.Itoa(port))

}
