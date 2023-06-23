package slices

import "fmt"

func Demo2() {
	cities := []string{"ankara", "istanbul", "izmir"} //making slices

	fmt.Println(cities)
	fmt.Println(len(cities))
	citiescopy := make([]string, len(cities))
	fmt.Println(citiescopy)
	fmt.Println(len(citiescopy))
	copy(citiescopy, cities) //to copy one slice to other
	fmt.Println(citiescopy)
	fmt.Println(len(citiescopy))

	cities = append(cities, "adana") //adding component to slice
	fmt.Println(cities)
	fmt.Println(len(cities))
	fmt.Println(citiescopy)
	fmt.Println(len(citiescopy))

	fmt.Println(cities[1:3]) //divide slice from 1st indis to 3rd indis.this fonction include first component but does not include last component.

}
