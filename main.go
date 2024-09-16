package main

import (
	"fmt"
)

type User struct {
	name  string
	score int
}

var rockPaperScissors = map[int]string{1: "rock", 2: "paper", 3: "scissors"}

var users []User

func create_new_player() {
	var name string

	fmt.Println("Please enter your name : ")

	_, err := fmt.Scan(&name)

	if err != nil {
		return
	}

	if !isValidName(name) {
		fmt.Println("The name must contain at least one character ")
		create_new_player()

	}

	if userExists(name) {
		fmt.Println("User with this name already exists.")
		create_new_player()
	}

	var user = User{name: name, score: 0}

	users = append(users, user)
}

func main() {

}
