package games

import (
	"fmt"
	"math/rand"
	"time"
)

func TwoGames() {

	gamenumber := 0
	fmt.Print(`
		
			
			██     ██ ███████ ██       ██████  ██████  ███    ███ ███████     ████████  ██████       ██████   █████  ███    ███ ███████ ███████ 
			██     ██ ██      ██      ██      ██    ██ ████  ████ ██             ██    ██    ██     ██       ██   ██ ████  ████ ██      ██      
			██  █  ██ █████   ██      ██      ██    ██ ██ ████ ██ █████          ██    ██    ██     ██   ███ ███████ ██ ████ ██ █████   ███████ 
			██ ███ ██ ██      ██      ██      ██    ██ ██  ██  ██ ██             ██    ██    ██     ██    ██ ██   ██ ██  ██  ██ ██           ██ 
			 ███ ███  ███████ ███████  ██████  ██████  ██      ██ ███████        ██     ██████       ██████  ██   ██ ██      ██ ███████ ███████ 	
			
				   `)
	fmt.Println("for math game press '1' for guess game press '2'")
	fmt.Scanln(&gamenumber)

	if gamenumber == 1 {

		number := 0
		var sth int = 0

		fmt.Print("keep a number in your mind and dont tell me what you kept. ")
		fmt.Scanln(&number)
		fmt.Print("add +1 to your number")
		fmt.Scanln(&number)
		fmt.Print("double new number")
		fmt.Scanln(&number)
		fmt.Print("add +4 to new number")
		fmt.Scanln(&number)
		fmt.Print("divide by 2 new number")
		fmt.Scanln(&number)
		fmt.Print("substract your initial value from the number you get lastly.")
		fmt.Scanln(&number)
		fmt.Print("Was your number 3? If it is true put '1' if it is not put '2'")
		fmt.Scanln(&sth)

		if sth == 1 {
			fmt.Println("end of the game")

		} else if sth == 2 {
			fmt.Println("please check your calculations :d")
		}

		for sth != 1 && sth != 2 {

			fmt.Print("you can only press 1 and 2")
			fmt.Scanln(&sth)

			if sth == 1 {
				fmt.Print("end of the game")

			} else if sth == 2 {
				fmt.Print("please check your calculations :d")
			}

		}

	} else if gamenumber == 2 {

		rand.Seed(time.Now().UnixNano())
		myNumber := rand.Intn(100) + 1 // use a random number
		estimatedNumber := 1
		guessAmount := 0

		fmt.Println("try to estimate the number that I choose randomly.")
		fmt.Scanln(&estimatedNumber)
		guessAmount = guessAmount + 1

		if estimatedNumber > 0 && estimatedNumber <= 100 {

			for estimatedNumber != myNumber {
				if estimatedNumber < myNumber {
					fmt.Println("make it bigger ;)")
					fmt.Scanln(&estimatedNumber)
					guessAmount = guessAmount + 1

				} else {
					fmt.Println("make it samller :D")
					fmt.Scanln(&estimatedNumber)
					guessAmount = guessAmount + 1

				}
			}

			guessStanding := " "

			if guessAmount > 0 && guessAmount <= 3 {
				guessStanding = "mind reader"

			} else if guessAmount > 3 && guessAmount <= 5 {
				guessStanding = "good job"
			} else {
				guessStanding = "quit from that job"
			}

			fmt.Printf("congratulations. %v guess : %v ", guessAmount, guessStanding)

		} else {
			fmt.Println("please read question correctly")

		}

	}
}
