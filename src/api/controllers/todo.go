package handlers

import (
	"encoding/json"
	"net/http"
	"todo-v1/src/api/models"
	"todo-v1/src/api/services"
)

func ListTodos(writer http.ResponseWriter, request *http.Request) {
	response := services.ListTodos()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}

func CreateTodo(writer http.ResponseWriter, request *http.Request) {
	var input models.CreateTodoInput
	json.NewDecoder(request.Body).Decode(&input)

	response := services.CreateTodo(input)

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
