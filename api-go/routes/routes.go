package routes

import (
	"collections/pkg/auth"
	"collections/pkg/collection"
	"collections/pkg/movie"
	"collections/pkg/user"

	"github.com/gin-gonic/gin"
)

// CreateRoutes create application routes
func CreateRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	r.POST("/register", user.CreateUser)
	r.POST("/login", user.Login)

	userRoutes := r.Group("/user").Use(auth.CheckAccessToken())
	userRoutes.GET("/:id", user.GetUser)

	// Collections
	collectionRoutes := r.Group("/collection").Use(auth.CheckAccessToken())
	collectionRoutes.POST("/create", collection.CreateCollection)
	collectionRoutes.POST("/add/movie/", collection.AddToCollection)
	collectionRoutes.GET("/:id", collection.GetCollection)

	collectionsRoutes := r.Group("/collections").Use(auth.CheckAccessToken())
	collectionsRoutes.GET("/self", collection.GetMyCollections)

	// Movies
	movieRoutes := r.Group("/movie").Use(auth.CheckAccessToken())
	movieRoutes.POST("/create", movie.CreateMovie)
	movieRoutes.GET("/search/:searchterm", movie.SearchMovie)
	return r
}
