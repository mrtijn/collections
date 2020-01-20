package auth

import "github.com/dgrijalva/jwt-go"

// Token model
type Token struct {
	UserID uint
	Email  string
	*jwt.StandardClaims
}
