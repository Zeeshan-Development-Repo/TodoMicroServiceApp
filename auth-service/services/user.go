package services

import (
	"auth-service/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUser(email string) (models.User, error) {
	filter := bson.M{"email": email}
	var existingUser models.User
	err := models.UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)

	return existingUser, err
}

func CreateUser(username string, email string, password string) error {
	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Email:     email,
		Password:  password,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := models.UserCollection.InsertOne(context.TODO(), newUser)
	return err
}
