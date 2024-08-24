package routes

import (
	"blackjack-app/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// Authentication routes
	router.HandleFunc("/register", controllers.Register).Methods("GET", "POST")
	router.HandleFunc("/login", controllers.Login).Methods("GET", "POST")

	// Protected route
	router.HandleFunc("/game", controllers.Game).Methods("GET")

	return router
}
