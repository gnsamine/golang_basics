package arrays

import "fmt"

func Demo3() {
	numbers := [6]int{16, 25, 36, 39, 45, 35}
	fmt.Println(numbers) // to write as an array

	for i := 0; i < len(numbers); i++ { // using len function we can see amount of
		fmt.Println(numbers[i]) //to wirte all numbers bottom to bototm

	}

	biggestValue := numbers[0]

	for i := 0; i < len(numbers); i++ {

		if numbers[i] > biggestValue {
			biggestValue = numbers[i]
		}

	}
	fmt.Println("biggest value:", biggestValue)

}
