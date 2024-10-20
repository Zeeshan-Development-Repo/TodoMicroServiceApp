package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Provider struct {
	Google struct {
		ID string `bson:"id" json:"id"`
	} `bson:"google,omitempty" json:"google,omitempty"`

	Facebook struct {
		ID string `bson:"id" json:"id"`
	} `bson:"facebook,omitempty" json:"facebook,omitempty"`
}

// User struct to represent a user
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username  string             `bson:"username" json:"username"`
	Email     string             `bson:"email" json:"email"`
	Name      string             `bson:"name" json:"name"`
	Password  string             `bson:"password" json:"password"`
	Providers Provider           `bson:"providers" json:"providers"`
	CreatedAt primitive.DateTime `bson:"created_at" json:"created_at"`
}

var UserCollection *mongo.Collection
