package pointers

import (
	"fmt"
)

// a pointer is a variable that holds the memory address of another variable.
// It allows you to indirectly access and modify the value of that variable.
// Pointers are useful when you want to pass a reference to a variable or manipulate data directly in memory.
func Demo1(num int) {
	num = num + 1
	fmt.Println("edited number:", num)

}
