package deferstatement

import "fmt"

func Demo2(number int32) string {
	defer DeferFunc()
	if number%2 == 0 {
		return "even nummber"

	}

	if number%2 != 0 {
		return "odd number"
	}

	return " "
}

func Test() {
	result := Demo2(7)
	fmt.Println(result)
}

func DeferFunc() {
	fmt.Println("finish")

}
