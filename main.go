package main

import "fmt"

type user struct {
	name  string
	score int
}

var rockPaperScissors = map[int]string{1: "rock", 2: "paper", 3: "scissors"}

var scores = map[struct{}]int{}

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

func main() {

}
