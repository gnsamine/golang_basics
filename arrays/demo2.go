package arrays

import "fmt"

func Demo2() {

	var cities [5]string

	cities[0] = "ankara"
	cities[1] = "istanbul"
	cities[2] = "muğla"
	cities[3] = "kastamonu"
	cities[4] = "manisa"
	fmt.Println(cities)

	for i := 0; i < 5; i++ {
		fmt.Println(cities[i]) // if you want to see all cities in different lines
	}

}
