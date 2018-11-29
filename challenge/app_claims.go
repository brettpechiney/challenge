package challenge

import "github.com/dgrijalva/jwt-go"

// AppClaims represent the set of valid JWT claims for the application.
type AppClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
}
