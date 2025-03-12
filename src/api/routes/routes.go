package routes

import (
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	RegisterHealthCheckRoutes(router)
	RegisterTodoRoutes(router)

	return router
}
