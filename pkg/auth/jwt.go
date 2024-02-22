package auth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// jwtKey is the secret key used for generating JWTs
var jwtKey = []byte("supersecretkey")

// JWTClaim represents the JWT Claims which will become the payload of the JWT
type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a JWT string using the provided email and username
func GenerateJWT(email string, username string) (tokenString string, err error) {
	// Set expiration time to 1 hour from now
	expirationTime := time.Now().Add(1 * time.Hour)
	// Create JWT claims
	claims := &JWTClaim{
		Email:    email,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Create JWT token with claims and signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate signed token string
	tokenString, err = token.SignedString(jwtKey)
	return
}

// ValidateToken validates the provided JWT token
func ValidateToken(signedToken string) (err error) {
	// Parse the token with claims
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}

	// Extract claims from token
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	// Check if token is expired
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	}

	return
}
