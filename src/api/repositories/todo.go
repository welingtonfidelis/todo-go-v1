package repositories

import (
	"fmt"
	"time"
	"todo-v1/src/api/models"
)

// TODO use database and repositories instead of memory array
var description = "description A"
var todoList = []models.Todo{
	{
		ID:          1,
		Name:        "test A",
		Status:      "todo",
		Date:        time.Date(2025, time.March, 12, 20, 6, 17, 0, time.UTC),
		Description: &description,
	},
	{
		ID:     2,
		Name:   "test B",
		Status: "todo",
		Date:   time.Date(2025, time.March, 12, 20, 6, 17, 0, time.UTC),
	},
}

func ListTodos() []models.Todo {
	return todoList
}

func CreateTodo(newTodoInput models.CreateTodoInput) models.Todo {
	fmt.Printf("%+v\n", newTodoInput)

	newId := len(todoList) + 1
	newTodo := models.Todo{
		ID:          newId,
		Name:        newTodoInput.Name,
		Description: newTodoInput.Description,
		Status:      newTodoInput.Status,
		Date:        newTodoInput.Date,
	}

	todoList = append(todoList, newTodo)

	return newTodo
}
