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
		selected, err := mainMenu()

		if err != nil {
			log.Fatal("Failed to chose main menu option", err)
			return
		}

		if selected == 1 {

			selectedInStartMenu, err := startGameMenu()

			if err != nil {
				log.Fatal("Failed to chose start game  menu option", err)
				return
			}

			if selectedInStartMenu == 1 {
				err = gameAi()

				if err != nil {
					log.Fatal("something failed in game: ", err)
					return
				}

				fmt.Println("Please enter anything to redirect... ")
				_, err := fmt.Scanln()
				if err != nil {
					return
				}

			} else if selectedInStartMenu == 2 {
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
			err := leaderboardMenu()
			if err != nil {
				log.Fatal("Something failed in leaderboard menu", err)
				return
			}

			fmt.Println("Please select anything to redirect ...")
			_, err = fmt.Scanln()
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
