package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func getWinner(playerNumber, aiNumber int) int {
	var winCondition = map[int]map[int]int{
		1: {2: 2, 3: 1, 1: 0},
		2: {1: 2, 3: 3, 2: 0},
		3: {1: 1, 2: 3, 3: 0},
	}

	return winCondition[playerNumber][aiNumber]
}

func gameAi() error {
	player, err := selectPlayer()

	if err != nil {
		return err
	}

	if player == nil {
		err = errors.New("please create a player first")
		return err
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
			return err
		}

		if playerNumber < 1 || playerNumber > 3 || aiNumber < 1 || aiNumber > 3 {
			fmt.Println("Invalid choice. Please choose 1, 2 or 3")
			continue
		}

		winner := getWinner(playerNumber, aiNumber)

		if winner == playerNumber {
			playerScore++
			fmt.Println("Congratulations you won this round !")

		} else if winner == aiNumber {
			aiScore++
			fmt.Println("You lost... !")

		} else {
			fmt.Println("Draw ... !")
		}
		fmt.Print("Wanna play another match: (yes-no)")
		var isRetry string

		_, err = fmt.Scan(&isRetry)
		if err != nil {
			return err
		}

		if strings.ToLower(isRetry) != "yes" {
			break
		}

	}

	if playerScore > aiScore {
		fmt.Println("You won totally")
		player.Score++
		err = userRepository.UpdateUser(player)

		if err != nil {
			return err
		}

	} else if playerScore < aiScore {
		fmt.Println("You lost totally")
	} else {
		fmt.Println("Draw... !")
	}

	return nil
}

func gamePlayer() error {
	player1, err := selectPlayer()

	if err != nil {
		return err
	}

	player2, err := selectPlayer()

	if err != nil {
		return err
	}

	var player1Score = 0
	var player2Score = 0

	for player1 == player2 {
		player2, err = selectPlayer()

		if err != nil {
			return err
		}
	}

	for {
		fmt.Printf("The game will be start with %v and %v\n", player1.Name, player2.Name)
		fmt.Println("The game will be start in 3s ...")
		time.Sleep(3)

		fmt.Println("1 - ROCK")
		fmt.Println("2 - PAPER")
		fmt.Println("3 - SCISSORS")
		var player1Choice int
		var player2Choice int

		fmt.Printf("Whats your choice %v : \n", player1.Name)
		_, err := fmt.Scan(&player1Choice)
		if err != nil {
			return err
		}

		fmt.Printf("Whats your choice %v : \n", player2.Name)

		_, err = fmt.Scan(&player2Choice)

		if err != nil {
			return err
		}

		if player1Choice < 1 || player1Choice > 3 || player2Choice < 1 || player2Choice > 3 {
			fmt.Println("Invalid choice. Please choose 1, 2 or 3")
			continue
		}

		winner := getWinner(player1Choice, player2Choice)

		if winner == player1Choice {
			fmt.Printf("Player %v won the round\n", player1.Name)
			player1Score++
		} else if winner == player2Choice {
			fmt.Printf("Player %v won the round\n", player2.Name)
			player2Score++
		} else {
			fmt.Println("Draw ... !")
		}

		fmt.Print("Wanna play another match: (yes-no)")
		var isRetry string

		_, err = fmt.Scan(&isRetry)
		if err != nil {
			return err
		}

		if strings.ToLower(isRetry) != "yes" {
			break
		}

	}

	if player1Score > player2Score {
		fmt.Printf("Player %v won totally\n", player1.Name)
		player1.Score++
		err = userRepository.UpdateUser(player1)

		if err != nil {
			return err
		}

	} else if player1Score < player2Score {
		fmt.Printf("Player %v won totally\n", player2.Name)
		player2.Score++
		err = userRepository.UpdateUser(player2)

		if err != nil {
			return err
		}

	} else {
		fmt.Println("Draw ... !")
	}
	return nil
}
