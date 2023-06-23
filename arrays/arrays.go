package arrays

import "fmt"

func Demo1() {
	// indexs begin with zeroth index
	var numbers [6]int //lets create array contains six term
	numbers[2] = 45    // lets second index be 45
	fmt.Println(numbers)

	fmt.Println(numbers[2]) // lets just write second index

}
