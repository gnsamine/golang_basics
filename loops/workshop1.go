package loops

import "fmt"

//
func Demo3() {
	myNumber := 15
	estimatedNumber := 1
	guessAmount := 0

	fmt.Println("keep a number in your mind between 1 to 100")
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
