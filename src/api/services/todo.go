package services

import (
	"todo-v1/src/api/models"
	"todo-v1/src/api/repositories"
)

func ListTodos() []models.Todo {
	return repositories.ListTodos()
}

func CreateTodo(newTodo models.CreateTodoInput) models.Todo {
	return repositories.CreateTodo(newTodo)
}
