package jwt_utils

import (
	"auth-service/models"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/shareed2k/goth_fiber"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var jwtSecret = []byte("your-secret-key") // Replace with a secure key

func generateJWT(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    user.Email,
		"name":     user.Name,
		"username": user.Username,  // Include Username in JWT claims
		"provider": user.Providers, // Change provider to Providers object
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func HandleOAuthCallback(c *fiber.Ctx) error {
	user, err := goth_fiber.CompleteUserAuth(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Authentication failed")
	}

	// Check if user already exists in MongoDB
	filter := bson.M{"email": user.Email}
	var existingUser models.User
	err = models.UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)

	if err == mongo.ErrNoDocuments {
		// If user doesn't exist, insert new user
		newUser := models.User{
			ID:        primitive.NewObjectID(),
			Username:  user.Name,
			Email:     user.Email,
			Name:      user.Name,
			CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		}
		_, err := models.UserCollection.InsertOne(context.TODO(), newUser)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Error inserting user in MongoDB")
		}

		fmt.Println("New user inserted:", newUser)
	} else if err != nil {
		return err
		return c.Status(fiber.StatusInternalServerError).SendString("Error checking user in MongoDB")
	} else {
		fmt.Println("User already exists:", existingUser)
	}

	// Normalize the existing user for JWT token generation
	normalizedUser := existingUser // Use the existing user directly

	token, err := generateJWT(normalizedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not generate JWT")
	}

	data := c.JSON(fiber.Map{
		"token": token, // Return JWT to client
		"user":  normalizedUser,
	})

	return data
}
