package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rockPaperScissors = map[int]string{1: "rock", 2: "paper", 3: "scissors"}

func getWinner(playerNumber, aiNumber int) int {

}

func gameAi() {
	player := selectPlayer()

	if player == nil {
		fmt.Println("Create a player first")
		return
	}

	var aiScore = 0
	var playerScore = 0

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for {
		var aiNumber = rng.Intn(3) + 1
		fmt.Println("1 - ROCK")
		fmt.Println("2 - PAPER")
		fmt.Println("3 - SCISSORS")
		fmt.Println("Which do you chose: ")
		var playerNumber int

		_, err := fmt.Scan(&playerNumber)
		if err != nil {
			return
		}

		if playerNumber < 1 || playerNumber > 3 || aiNumber < 1 || aiNumber > 3 {
			fmt.Println("Invalid choice. Please choose 1, 2 or 3")
			continue
		}

		winner := getWinner(playerNumber, aiNumber)
	}
}

func gamePlayer() {

}
