package main

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
)

type User struct {
	name  string
	score int
}

var rockPaperScissors = map[int]string{1: "rock", 2: "paper", 3: "scissors"}

var users []User

func main_menu() int {
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
		return 0
	}

	return option
}

func start_game_menu() int {
	fmt.Println("Welcome---")
	fmt.Println("Please Enter your choice: ")
	fmt.Println("1 - Play vs ai")
	fmt.Println("2 - Play vs Player")
	fmt.Println("3 - Back to main menu")
	var option int

	_, err := fmt.Scan(&option)

	if err != nil {
		return 0
	}

	return option
}

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

func leaderboard_menu() {
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Name", "Score"})
	for _, user := range users {
		table.Append([]string{user.name, fmt.Sprintf("%v", user.score)})
	}

	table.Render()
}
func main() {

}
