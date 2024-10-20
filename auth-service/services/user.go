package services

import (
	"auth-service/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckUser(username string) (models.User, error) {
	filter := bson.M{"username": username}
	var existingUser models.User
	err := models.UserCollection.FindOne(context.TODO(), filter).Decode(&existingUser)

	return existingUser, err
}

func CreateUser(username string, password string) error {
	newUser := models.User{
		ID:        primitive.NewObjectID(),
		Username:  username,
		Password:  password,
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
	}

	_, err := models.UserCollection.InsertOne(context.TODO(), newUser)
	return err
}
