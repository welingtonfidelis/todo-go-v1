package services

import (
	"time"
	"todo-v1/src/api/models"
)

// TODO use database and repositories instead of memory array
var description = "description A"
var todoList = []models.Todo{
	{
		Name:        "test A",
		Status:      "todo",
		Date:        time.Date(2025, time.March, 12, 20, 6, 17, 0, time.UTC),
		Description: &description,
	},
	{
		Name:   "test B",
		Status: "todo",
		Date:   time.Date(2025, time.March, 12, 20, 6, 17, 0, time.UTC),
	},
}

func ListTodos() []models.Todo {
	return todoList
}
