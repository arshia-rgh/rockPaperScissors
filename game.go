package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getWinner(playerNumber, aiNumber int) int {
	var winCondition = map[int]map[int]int{
		1: {2: 2, 3: 1},
		2: {1: 2, 3: 3},
		3: {1: 1, 2: 3},
	}

	return winCondition[playerNumber][aiNumber]
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

		if winner == playerNumber {
			playerScore++
			fmt.Println("Congratulations you won this round !")

		} else {
			aiScore++
			fmt.Println("You lost... !")

		}
		fmt.Print("Wanna play another match: (yes-no)")
		var isRetry string

		_, err = fmt.Scan(&isRetry)
		if err != nil {
			return
		}

		if strings.ToLower(isRetry) != "yes" {
			break
		}

	}

	if playerScore > aiScore {
		fmt.Println("You won totally")
		player.score++
	} else if playerScore < aiScore {
		fmt.Println("You lost totally")
	} else {
		fmt.Println("Draw... !")
	}
}

func gamePlayer() {

}
