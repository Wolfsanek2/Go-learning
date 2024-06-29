// Printing a user’s name based on user input
package main

import (
	"fmt"
	"os"
)

func createUsers() map[string]string {
	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}
	return users
}

func getUser(id string) (string, bool) {
	users := createUsers()
	user, exists := users[id]
	return user, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Enter user name")
		os.Exit(1)
	}
	id := os.Args[1]
	user, exists := getUser(id)
	if !exists {
		fmt.Println("User not found")
		os.Exit(1)
	}
	fmt.Println(user)
}
