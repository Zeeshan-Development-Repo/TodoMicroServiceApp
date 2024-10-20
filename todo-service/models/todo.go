package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Todo represents a todo item in the database
type Todo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	UserId      string             `bson:"user_id" json:"user_id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Completed   bool               `bson:"completed" json:"completed"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// TodoCollection represents the MongoDB collection
var TodoCollection *mongo.Collection
