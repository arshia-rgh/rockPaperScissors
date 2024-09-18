package main

import "fmt"

func main() {
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
			createNewPlayer()
			fmt.Println("Please select anything to redirect ...")
			_, err := fmt.Scanln()
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
