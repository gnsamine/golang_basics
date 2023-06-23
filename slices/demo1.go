package slices

import "fmt"

func Demo1() {

	names := make([]string, 3) // making slice

	names[0] = "amin"
	names[1] = "sd"
	names[2] = "eso"

	fmt.Println(names)
	fmt.Println(len(names))

	names = append(names, "lol") //when you want to add component to your list
	fmt.Println(names)
	fmt.Println(len(names))
}
