package routes

import (
	handlers "todo-v1/src/api/controllers"

	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(router *mux.Router) {
	todoRouter := router.PathPrefix("/todos").Subrouter()

	todoRouter.HandleFunc("", handlers.ListTodos).Methods("GET")
	todoRouter.HandleFunc("", handlers.CreateTodo).Methods("POST")
	// todoRouter.HandleFunc("/{id}", handlers.GetTodo).Methods("GET")
	// todoRouter.HandleFunc("/{id}", handlers.UpdateTodo).Methods("PUT")
	// todoRouter.HandleFunc("/{id}", handlers.DeleteTodo).Methods("DELETE")
}
