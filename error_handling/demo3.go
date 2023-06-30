package errorhandling

import (
	"fmt"
)

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string {
	// error is a methode for BorderException
	//when we do this we say : for BorderException struct, there is a methode named "error"

	return fmt.Sprintf("%d,---%s", b.parameter, b.message)

}

func Guess2(guess int) (string, error) {
	if guess < 1 || guess > 100 {
		return "", &BorderException{guess, "Out of the borders"}
	}

	return "bildiniz", nil

}

// you will write in main go "fmt.Println(errorhandling.Guess2(150))"
