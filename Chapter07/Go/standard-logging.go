package main

import (
	"log"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("Failed to get user with error: %v", err)
	}

	log.Printf("Current user is %s", user.Username)
}
