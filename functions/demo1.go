package functions

import "fmt"

func Add(number1 int, number2 int, number3 int) {
	// this method not commonly used.
	result := number1 + number2 + number3
	fmt.Println("result: ", result)

}

func Add2(number1 int, number2 int) int {
	result := number1 + number2
	return result
}

func SayHi2(username string) {
	fmt.Println("user name:", username)

}
func SayHi1(username string) {
	fmt.Println(username)
	fmt.Scanln(&username)
	fmt.Println("hi", username)
}

func SayHi() {
	fmt.Println("amin")

}
