package goroutines

import (
	"fmt"
	"time"
)

func EvenNumbers() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even Number:", i)
		time.Sleep(time.Second)

	}

}

func OddNumbers() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd Number:", i)
		time.Sleep(time.Second)

	}

}
