package jwt_service

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("I Want Something Just Like this Todo do do do dooo!")

// CreateToken generates a JWT token with custom claims
func CreateToken(email string, name string, username string) (string, error) {
	claims := jwt.MapClaims{
		"email":    email,
		"name":     name,
		"username": username,                              // Include Username in JWT claims
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time
	}

	// Create token using the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token using the secret key
	return token.SignedString(secretKey)
}

// VerifyToken parses and validates the given token
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	// Check if the token is valid
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
