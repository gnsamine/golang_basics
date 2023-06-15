package loops

import (
	"fmt"
)

func Game1() {

	number := 0
	var sth int = 0

	fmt.Println("keep a number in your mind and dont tell me what you kept. you can give me every number that you want to fool me for every step")
	fmt.Scanln(&number)
	fmt.Println("add +1 to your number")
	fmt.Scanln(&number)
	fmt.Println("double new number")
	fmt.Scanln(&number)
	fmt.Println("add +4 to new number")
	fmt.Scanln(&number)
	fmt.Println("divide by 2 new number")
	fmt.Scanln(&number)
	fmt.Println("substract your initial value from the number you get lastly.")
	fmt.Scanln(&number)
	fmt.Println("Was your number 3? If it is true put '1' if it is not put '2'")
	fmt.Scanln(&sth)

	if sth == 1 {
		fmt.Println("end of the game")

	} else if sth == 2 {
		fmt.Println("please check your calculations :d")
	}

	for sth != 1 && sth != 2 {

		fmt.Println("you can only press 1 and 2")
		fmt.Scanln(&sth)

		if sth == 1 {
			fmt.Println("end of the game")

		} else if sth == 2 {
			fmt.Println("please check your calculations :d")
		}

	}
}
