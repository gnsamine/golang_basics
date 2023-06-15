package loops

import "fmt"

//
func Demo3() {
	myNumber := 10
	estimatedNumber := 1

	fmt.Println("keep a number in your mind between 10 to 30")
	fmt.Scanln(&estimatedNumber)

	if estimatedNumber == myNumber {
		fmt.Println("congratulation. Now you are master of reading mind.")
	} else if estimatedNumber < myNumber {
		fmt.Println("make it bigger ;)")

	} else {
		fmt.Println("make it samller :D")
	}

}
