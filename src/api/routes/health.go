package routes

import (
	handlers "todo-v1/src/api/controllers"

	"github.com/gorilla/mux"
)

func RegisterHealthCheckRoutes(router *mux.Router) {
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")
}
