package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"blackjack-app/models"
)

/* func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Parse form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Register user in the database
	err := models.CreateUser(username, password)
	if err != nil {
		http.Error(w, "Unable to register user", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "User registered successfully")
} */

func Register(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is GET (to display the form)
	if r.Method == "GET" {
		// Parse the template and render it
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
		return
	}

	// For POST request: handle form submission
	if r.Method == "POST" {
		// Parse form values
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Register user in the database
		err := models.CreateUser(username, password)
		if err != nil {
			http.Error(w, "Unable to register user", http.StatusBadRequest)
			return
		}

		// Send a success message back
		fmt.Fprintf(w, "User registered successfully")
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl := template.Must(template.ParseFiles("templates/login.html"))
		tmpl.Execute(w, nil)
		return
	}

	// Parse form data
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authenticate user
	token, err := models.AuthenticateUser(username, password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return JWT token
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func Game(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the Blackjack game!"))
}
