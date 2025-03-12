package handlers

import (
	"encoding/json"
	"net/http"
	"todo-v1/src/api/services"
)

func ListTodos(writer http.ResponseWriter, request *http.Request) {
	response := services.ListTodos()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
