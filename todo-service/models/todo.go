package models

import "errors"

// Todo struct represents a todo item
type Todo struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Complete bool   `json:"complete"`
}

// In-memory todo data
var todos = []Todo{
	{ID: "1", Title: "Buy groceries", Complete: false},
	{ID: "2", Title: "Learn Go Fiber", Complete: false},
}

// Get all todos
func GetAllTodos() []Todo {
	return todos
}

// Get todo by ID
func GetTodoByID(id string) (Todo, error) {
	for _, todo := range todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return Todo{}, errors.New("Todo not found")
}

// Add a new todo
func AddTodo(todo *Todo) {
	todos = append(todos, *todo)
}

// Update todo by ID
func UpdateTodo(id string, updatedTodo *Todo) (Todo, error) {
	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Complete = updatedTodo.Complete
			return todos[i], nil
		}
	}
	return Todo{}, errors.New("Todo not found")
}

// Delete todo by ID
func DeleteTodo(id string) error {
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return nil
		}
	}
	return errors.New("Todo not found")
}
