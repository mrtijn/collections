package auth

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword bcrypts password
func HashPassword(password string) ([]byte, error) {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return pass, err
}

// VerifyPassword verifies the hashed pass
func VerifyPassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err
}

// CheckAccessToken verify token
func CheckAccessToken() gin.HandlerFunc {
	return func(c *gin.Context) {

		var tokenString = c.GetHeader("x-access-token")
		tokenString = strings.TrimSpace(tokenString)

		if tokenString == "" {
			c.JSON(http.StatusForbidden, "No token send")
			c.Abort()
			return
		}

		tk := &Token{}

		token, err := jwt.ParseWithClaims(tokenString, tk, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("jwtSecret")), nil
		})

		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusForbidden, "Token expired")
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*Token)
		tokenExpired := claims.StandardClaims.ExpiresAt <= time.Now().Unix()

		if ok && token.Valid && tokenExpired {
			c.JSON(http.StatusForbidden, "Token expired")
			c.Abort()
			return
		}

		// If all ok, update session
		session := sessions.Default(c)
		session.Set("userID", claims.UserID)
		session.Save()

		c.Next()

	}
}
