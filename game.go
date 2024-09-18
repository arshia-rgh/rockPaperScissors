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
	player1 := selectPlayer()
	player2 := selectPlayer()
	var player1Score = 0
	var player2Score = 0

	for player1 == player2 {
		player2 = selectPlayer()
	}

	for {
		fmt.Printf("The game will be start with %v and %v\n", player1.name, player2.name)
		fmt.Println("The game will be start in 3s ...")
		time.Sleep(3)

		fmt.Println("1 - ROCK")
		fmt.Println("2 - PAPER")
		fmt.Println("3 - SCISSORS")
		var player1Choice int
		var player2Choice int

		fmt.Printf("Whats your choice %v : \n", player1.name)
		_, err := fmt.Scan(&player1Choice)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Whats your choice %v : \n", player2.name)

		_, err = fmt.Scan(&player2Choice)

		if err != nil {
			fmt.Println(err)
			return
		}

		if player1Choice < 1 || player1Choice > 3 || player2Choice < 1 || player2Choice > 3 {
			fmt.Println("Invalid choice. Please choose 1, 2 or 3")
			continue
		}

		winner := getWinner(player1Choice, player2Choice)

		if winner == player1Choice {
			fmt.Printf("Player %v won the round", player1.name)
			player1Score ++
		} else
	}
}
