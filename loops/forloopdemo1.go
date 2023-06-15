package loops

import "fmt"

func Demo1() {
	var sth string = "hi"
	i := 1

	for i <= 5 {
		fmt.Println(sth)
		i = i + 2
		//if you don't give new value to "i" you get infinite loop.

	}
}
