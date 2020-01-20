package routes

import (
	"collections/pkg/auth"
	"collections/pkg/user"
	"github.com/gin-gonic/gin"
)

// CreateRoutes create application routes
func CreateRoutes(r *gin.RouterGroup) *gin.RouterGroup {

	r.POST("/register", user.CreateUser)
	r.POST("/login", user.Login)

	userRoutes := r.Group("/user").Use(auth.CheckAccessToken())
	userRoutes.GET("/:id", user.GetUser)

	return r
}
