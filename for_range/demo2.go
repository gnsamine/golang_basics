package for_range

import "fmt"

// print addition of odd numbers between 1 and 10
func Demo2() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	addition := 0

	for i, number := range numbers {

		if number%2 == 1 {
			addition = addition + numbers[i]
		}

	}
	fmt.Println("addition:", addition)

}
