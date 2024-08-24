package models

import (
	"errors"

	"blackjack-app/database"
	"blackjack-app/utils"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	db := database.GetDB()
	_, err = db.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", username, hashedPassword)
	return err
}

func AuthenticateUser(username, password string) (string, error) {
	db := database.GetDB()
	var hashedPassword string
	err := db.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		return "", errors.New("user not found")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(username)
	return token, err
}
