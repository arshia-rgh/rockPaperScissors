package main

import (
	"errors"
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

	exists, err := userExists(name)

	if err != nil {
		return err
	}

	if exists {
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

func selectPlayer() (*models.User, error) {
	users, err := userRepository.GetAllUsers()

	if err != nil {
		return nil, err
	}

	fmt.Println("Please select the player you want to play as: ")

	if len(users) != 0 {
		for i, user := range users {
			fmt.Printf("%v - %v\n", i+1, user)
		}

		var chosenUser int
		fmt.Print("Chosen name: (Should enter the number) ")
		_, err := fmt.Scan(&chosenUser)
		if err != nil {
			return nil, err
		}

		if chosenUser >= 1 && chosenUser < len(users)+1 {
			return &users[chosenUser-1], nil
		}

		err = errors.New("invalid user selection")

		return nil, err

	}

	err = errors.New("there is no players registered")

	return nil, err
}
