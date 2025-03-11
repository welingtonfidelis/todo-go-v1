package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-v1/src/routes"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := routes.SetupRoutes()
	port := os.Getenv("PORT")
	serverAddress := ":" + port

	fmt.Println("Server running in port", port)
	log.Fatal(http.ListenAndServe(serverAddress, router))

}
