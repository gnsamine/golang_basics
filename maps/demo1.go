package maps

import "fmt"

func Demo1() {
	//key-value
	dictionary := make(map[string]string)
	dictionary["table"] = "masa"
	dictionary["book"] = "kitap"
	dictionary["pencil"] = "kalem"

	fmt.Println(dictionary)
	fmt.Println(dictionary["pencil"])
	fmt.Println("number of elements:", len(dictionary))

	delete(dictionary, "book")
	fmt.Println("number of elements:", len(dictionary))

	value, include := dictionary["book"]
	fmt.Println(value)
	fmt.Println("It is in the list: ", include)
}
