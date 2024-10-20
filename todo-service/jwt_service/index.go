package jwt_service

import (
	"fmt"
	"todo-service/models"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("I Want Something Just Like this Todo do do do dooo!")

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

// ExtractUserFromClaims extracts user information from the JWT claims
func ExtractUserFromClaims(claims jwt.MapClaims) (*models.User, error) {
	email, ok := claims["email"].(string)
	if !ok {
		fmt.Println("email claim is missing or not a string")
		return nil, fmt.Errorf("email claim is missing or not a string")
	}

	name, ok := claims["name"].(string)
	if !ok {
		fmt.Println("name claim is missing or not a string")
		return nil, fmt.Errorf("name claim is missing or not a string")
	}

	username, ok := claims["username"].(string)
	if !ok {
		fmt.Println("username claim is missing or not a string")
		return nil, fmt.Errorf("username claim is missing or not a string")
	}

	id, ok := claims["id"].(string)
	if !ok {
		fmt.Println("id claim is missing or not a string")
		return nil, fmt.Errorf("id claim is missing or not a string")
	}

	return &models.User{
		Id:       id,
		Email:    email,
		Name:     name,
		Username: username,
	}, nil
}
