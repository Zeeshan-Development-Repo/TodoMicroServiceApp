package models

// User struct to represent a user
type User struct {
	Id       string ` json:"id"`
	Username string `bson:"username" json:"username"`
	Email    string `bson:"email" json:"email"`
	Name     string `bson:"name" json:"name"`
}
