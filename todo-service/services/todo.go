package services

import (
	"todo-service/models"
	"todo-service/repos"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TodoService interface defines the methods for managing todos
type TodoService interface {
	CreateTodo(title string, description string, userId string) (*models.Todo, error)
	GetTodo(id primitive.ObjectID) (*models.Todo, error)
	UpdateTodo(id primitive.ObjectID, updatedTodo models.Todo) (*models.Todo, error)
	DeleteTodo(id primitive.ObjectID) error
	GetTodos(userId string) ([]models.Todo, error)
}

// todoService struct implements the TodoService interface
type todoService struct{}

// NewTodoService creates a new TodoService
func NewTodoService() TodoService {
	return &todoService{}
}

// CreateTodo creates a new todo
func (s *todoService) CreateTodo(title string, description string, userId string) (*models.Todo, error) {
	return repos.CreateTodo(title, description, userId)
}

// GetTodo retrieves a todo by ID
func (s *todoService) GetTodo(id primitive.ObjectID) (*models.Todo, error) {
	return repos.GetTodo(id)
}

// UpdateTodo updates an existing todo
func (s *todoService) UpdateTodo(id primitive.ObjectID, updatedTodo models.Todo) (*models.Todo, error) {
	return repos.UpdateTodo(id, updatedTodo)
}

// DeleteTodo deletes a todo by ID
func (s *todoService) DeleteTodo(id primitive.ObjectID) error {
	return repos.DeleteTodo(id)
}

// GetTodos retrieves all todos for a specific user
func (s *todoService) GetTodos(userId string) ([]models.Todo, error) {
	return repos.GetTodos(userId)
}
