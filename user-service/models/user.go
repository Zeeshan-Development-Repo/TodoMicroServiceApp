package models

import "errors"

// User struct to represent a user
type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// In-memory user data
var users = []User{
	{ID: "1", Name: "John Doe", Email: "john@example.com"},
	{ID: "2", Name: "Jane Doe", Email: "jane@example.com"},
}

// Get all users
func GetAllUsers() []User {
	return users
}

// Get user by ID
func GetUserByID(id string) (User, error) {
	for _, user := range users {
		if user.ID == id {
			return user, nil
		}
	}
	return User{}, errors.New("User not found")
}

// Add a new user
func AddUser(user *User) {
	users = append(users, *user)
}

// Update user by ID
func UpdateUser(id string, updatedUser *User) (User, error) {
	for i, user := range users {
		if user.ID == id {
			users[i].Name = updatedUser.Name
			users[i].Email = updatedUser.Email
			return users[i], nil
		}
	}
	return User{}, errors.New("User not found")
}

// Delete user by ID
func DeleteUser(id string) error {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return errors.New("User not found")
}
