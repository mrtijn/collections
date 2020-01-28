package user

import (
	"collections/database"
	"collections/models"
	"collections/pkg/auth"
	"collections/utils/formaterror"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var db = database.Connect()

// CreateUser create user in db
func CreateUser(c *gin.Context) {
	user := &models.User{}

	fmt.Println(c.Request.Body)

	err := c.Bind(user)
	// fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data"})
		return
	}

	pass, err := auth.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to encrypt password"})
		return
	}

	user.Password = string(pass)

	err = db.Create(user).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login function
func Login(c *gin.Context) {
	user := &models.User{}

	err := c.Bind(user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing data"})
		return
	}

	resp, err := getToken(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}

// GetUser gets user data
func GetUser(c *gin.Context) {
	user := &models.User{}

	var id = c.Param("id")

	err := db.Where("Id = ?", id).First(user).Error

	if err != nil {
		c.JSON(http.StatusNotFound, "User not found")
		return
	}

	c.JSON(http.StatusOK, user)
}

// getToken function
func getToken(email string, password string) (map[string]interface{}, error) {
	user := &models.User{}

	err := db.Where("Email = ?", email).First(user).Error

	if err != nil {
		return nil, formaterror.Parse("Could not find user")
	}

	err = auth.VerifyPassword(user.Password, password)
	fmt.Println(err)
	if err != nil {
		return nil, formaterror.Parse("wrong password")
	}

	expiresAt := time.Now().Add(time.Minute * 10000).Unix()
	tk := &auth.Token{
		UserID: user.ID,
		Email:  user.Email,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
		},
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)

	tokenString, err := token.SignedString([]byte(os.Getenv("jwtSecret")))

	if err != nil {
		return nil, formaterror.Parse("Could not get token")
	}

	var resp = map[string]interface{}{
		"token":     tokenString,
		"expiresAt": expiresAt,
	}
	return resp, nil
}
