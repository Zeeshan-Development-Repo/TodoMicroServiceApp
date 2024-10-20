package repos

import (
	"context"
	"time"
	"todo-service/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTodo inserts a new todo item into the database
func CreateTodo(title string, description string, userId string) (*models.Todo, error) {
	todo := models.Todo{
		ID:          primitive.NewObjectID(),
		Title:       title,
		Description: description,
		UserId:      userId,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	result, err := models.TodoCollection.InsertOne(context.Background(), todo)
	if err != nil {
		return nil, err
	}
	todo.ID = result.InsertedID.(primitive.ObjectID)
	return &todo, nil
}

// GetTodo retrieves a todo item by its ID
func GetTodo(id primitive.ObjectID) (*models.Todo, error) {
	var todo models.Todo
	err := models.TodoCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

// UpdateTodo updates an existing todo item
func UpdateTodo(id primitive.ObjectID, updatedTodo models.Todo) (*models.Todo, error) {
	updatedTodo.UpdatedAt = time.Now()
	_, err := models.TodoCollection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		bson.M{"$set": updatedTodo},
	)
	if err != nil {
		return nil, err
	}
	// Return the updated todo item
	return GetTodo(id)
}

// DeleteTodo removes a todo item from the database
func DeleteTodo(id primitive.ObjectID) error {
	_, err := models.TodoCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

// GetTodos retrieves all todo items for a specific user
func GetTodos(userId string) ([]models.Todo, error) {
	var todos []models.Todo
	cursor, err := models.TodoCollection.Find(context.Background(), bson.M{"user_id": userId})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	return todos, nil
}
