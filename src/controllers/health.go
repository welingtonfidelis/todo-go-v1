package handlers

import (
	"encoding/json"
	"net/http"
	"todo-v1/src/services"
)

func HealthCheck(writer http.ResponseWriter, request *http.Request) {
	response := services.HealthCheck()

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(response)
}
