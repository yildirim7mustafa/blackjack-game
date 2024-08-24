package main

import (
	"log"
	"net/http"

	"blackjack-app/database"
	"blackjack-app/routes"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Set up the routes
	router := routes.SetupRoutes()

	// Start the web server
	log.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", router)
}
