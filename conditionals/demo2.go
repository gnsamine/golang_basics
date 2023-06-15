package conditionals

import "fmt"

func Demo2() {
	var account float64 = 1000
	var moneytowithdrawn float64 = 980

	if moneytowithdrawn > account {
		fmt.Println("you have not sufficient money")

	} else {

		fmt.Println("machiene will give your money")
	}

}
func Demo3() {
	var account float64 = 1000
	var moneytowithdrawn float64 = 1000

	if moneytowithdrawn > account {
		fmt.Println("you have not sufficient money")
	} else if moneytowithdrawn == account {
		fmt.Println("warning:remaining amount of money= 0 ")
	} else {
		fmt.Println("machiene will give your money")
	}

}
