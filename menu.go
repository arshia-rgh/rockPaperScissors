package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

func mainMenu() (int, error) {
	fmt.Println("*****************************")
	fmt.Println("**********Welcome************")
	fmt.Println("Please enter what do you want to do: ")
	fmt.Println("1 - Start the game")
	fmt.Println("2 - Create new player")
	fmt.Println("3 - Check Leaderboard ")
	fmt.Println("4 - Exit...")
	var option int

	_, err := fmt.Scan(&option)
	if err != nil {
		return 0, err
	}

	return option, nil
}

func startGameMenu() (int, error) {
	fmt.Println("Welcome---")
	fmt.Println("Please Enter your choice: ")
	fmt.Println("1 - Play vs ai")
	fmt.Println("2 - Play vs Player")
	fmt.Println("3 - Back to main menu")
	var option int

	_, err := fmt.Scan(&option)

	if err != nil {
		return 0, err
	}

	return option, nil
}

func leaderboardMenu() error {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Name", "Score"})
	users, err := userRepository.GetAllUsers()

	if err != nil {
		return err
	}
	for _, user := range users {
		table.Append([]string{user.Name, fmt.Sprintf("%v", user.Score)})
	}

	table.Render()

	return nil
}
