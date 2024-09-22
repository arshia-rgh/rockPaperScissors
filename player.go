package main

import (
	"fmt"
	"rockPaperScissors/models"
)

func createNewPlayer() error {
	var name string

	fmt.Println("Please enter your name : ")

	_, err := fmt.Scan(&name)

	if err != nil {
		return err
	}

	if !isValidName(name) {
		fmt.Println("The name must contain at least one character ")
		err := createNewPlayer()
		if err != nil {
			return err
		}

	}

	if userExists(name) {
		fmt.Println("User with this name already exists.")
		err := createNewPlayer()
		if err != nil {
			return err
		}
	}

	user := models.NewUser(0, name, 0)

	err = userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil

}

func selectPlayer() *User {
	fmt.Println("Please select the player you want to play as: ")

	if len(users) != 0 {
		for i, user := range users {
			fmt.Printf("%v - %v\n", i+1, user)
		}

		var chosenUser int
		fmt.Print("Chosen name: (Should enter the number) ")
		_, err := fmt.Scan(&chosenUser)
		if err != nil {
			return nil
		}

		if chosenUser >= 1 && chosenUser < len(users)+1 {
			return &users[chosenUser-1]
		}

		fmt.Println("Invalid user selection.")
		return nil

	}

	fmt.Println("There is no players registered...")

	return nil
}
