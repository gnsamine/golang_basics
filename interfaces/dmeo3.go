package interfaces

import "fmt"

func verify(i interface{}) {
	number, ok := i.(int)
	// make type convertion and write true if it is number
	// if it is not ,swrite false and assign defoult value of int
	fmt.Println(number, ok)
}

func Demo3() {

	verify(5)
	verify("amin")

	var number2 interface{} = 9
	verify(number2)

	var number3 interface{} = "merhaba"
	verify(number3)
}
