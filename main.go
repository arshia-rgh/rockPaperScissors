package main

import (
	"fmt"
	"log"
	"rockPaperScissors/db"
	"rockPaperScissors/repositories"
)

var userRepository *repositories.UserRepository

func main() {
	err := db.ConnectDatabase()

	if err != nil {
		log.Fatal("Could not `connect` to the database", err)
	}

	err = db.MigrateDatabase()
	if err != nil {
		log.Fatal("Could not `migrate` to the database", err)
	}

	userRepository = repositories.NewUserRepository(db.DB)

	for {
		selected := mainMenu()

		if selected == 1 {

			selectedInStartMenue := startGameMenu()

			if selectedInStartMenue == 1 {
				gameAi()
				fmt.Println("Please enter anything to redirect... ")
				_, err := fmt.Scanln()
				if err != nil {
					return
				}

			} else if selectedInStartMenue == 2 {
				gamePlayer()
				fmt.Println("Please select anything to redirect ...")
				_, err := fmt.Scanln()
				if err != nil {
					return
				}

			} else {
				continue
			}

		} else if selected == 2 {
			err := createNewPlayer()
			if err != nil {
				log.Fatal("Could not create new user", err)
				return
			}
			fmt.Println("Please select anything to redirect ...")
			_, err = fmt.Scanln()
			if err != nil {
				return
			}
			continue

		} else if selected == 3 {
			leaderboardMenu()
			fmt.Println("Please select anything to redirect ...")
			_, err := fmt.Scanln()
			if err != nil {
				return
			}
			continue

		} else {
			fmt.Println("Goodbye...")
			break
		}
	}
}
