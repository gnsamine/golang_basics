package arrays

import "fmt"

func Demo4() {
	var numbers [3][2]int // identifying multivariable array

	numbers[0][0] = 1
	numbers[0][1] = 2

	numbers[1][0] = 3
	numbers[1][1] = 4

	numbers[2][0] = 5
	numbers[2][1] = 6

	for i := 0; i < 3; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(numbers[i][j])
			fmt.Print(" ")

		}
		fmt.Println()

	}

}
