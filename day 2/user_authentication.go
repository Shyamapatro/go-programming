// User Authentication:

// You are building an authentication system for an e-commerce platform. 
// Write a Go program that simulates user authentication by verifying usernames and passwords. 
// The program should store a list of registered users with their usernames and passwords.
// Users should be prompted to enter their username and password, and the program should validate the credentials.

package main

import (
	"fmt"
)

type User struct {
	Username string
	Password string
}

func AuthenticateUser(username, password string, users []User) bool {
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

func main() {
	// Registered users
	users := []User{
		{Username: "Shyama", Password: "12345"},
		{Username: "new", Password: "new"},
		{Username: "admin", Password: "admin"},
	}

	// Prompt the user to enter their username and password
	var username, password string
	fmt.Print("Enter your username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter your password: ")
	fmt.Scanln(&password)

	// Authenticate the user
	if AuthenticateUser(username, password, users) {
		fmt.Println("Authentication successful. You are logged in.")
	} else {
		fmt.Println("Authentication failed. Invalid username or password.")
	}
}
