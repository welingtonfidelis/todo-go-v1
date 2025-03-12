package routes

import (
	handlers "todo-v1/src/api/controllers"

	"github.com/gorilla/mux"
)

func RegisterTodoRoutes(router *mux.Router) {
	todoRouter := router.PathPrefix("/todos").Subrouter()

	todoRouter.HandleFunc("", handlers.ListTodos).Methods("GET")

	// todoRouter.Handle("", middlewares.ValidateTodo(http.HandlerFunc(handlers.CreateTodo))).Methods("POST")
	// todoRouter.HandleFunc("", handlers.GetTodo).Methods("POST")
	// todoRouter.HandleFunc("/{id}", handlers.GetTodo).Methods("GET")
	// todoRouter.HandleFunc("/{id}", handlers.UpdateTodo).Methods("PUT")
	// todoRouter.HandleFunc("/{id}", handlers.DeleteTodo).Methods("DELETE")
}
