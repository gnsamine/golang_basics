package loops

import "fmt"

//
func Demo3() {
	myNumber := 15
	estimatedNumber := 1

	fmt.Println("keep a number in your mind between 10 to 30")
	fmt.Scanln(&estimatedNumber)

	if estimatedNumber == myNumber {
		fmt.Println("congratulation. Now you are master of reading mind.")

	}
	for estimatedNumber != myNumber {
		if estimatedNumber < myNumber {
			fmt.Println("make it bigger ;)")
			fmt.Scanln(&estimatedNumber)

		} else {
			fmt.Println("make it samller :D")
			fmt.Scanln(&estimatedNumber)
		}
	}

}
