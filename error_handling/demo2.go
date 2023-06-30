package errorhandling

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func LetsGuess(guess int) (string, error) {

	rand.Seed(time.Now().UnixNano())
	myNumber := rand.Intn(100) + 1

	if guess < 1 || guess > 100 {
		return "", errors.New("choose a number between 1-100")
	}

	for guess <= 100 && guess >= 1 {
		if guess < myNumber {
			fmt.Println("make it bigger")
			fmt.Print("give me a number: ")
			fmt.Scanln(&guess)

		} else if guess > myNumber {
			fmt.Println("make it smaller")
			fmt.Print("give me a number: ")
			fmt.Scanln(&guess)
		} else {
			return "this is the number", nil
		}
	}

	return "give me a number between 1-100", errors.New("unexpected error occurred")
	//just in case there's a bug or issue in the loop logic,
	//and it keeps running indefinitely without reaching either
	//of the termination conditions,the line return "", errors.New("unexpected error occurred")
	//ensures that the function still returns an error, indicating that something unexpected happened.

}

func Demo2() {

	fmt.Print("Guess a number between 1 and 100: ")
	var yourNumber int
	fmt.Scanln(&yourNumber)

	result, err := LetsGuess(yourNumber)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(result)
	}
}
