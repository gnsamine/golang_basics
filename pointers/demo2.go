package pointers

import (
	"fmt"
	"unsafe"
)

func Demo2(numbers []int) {
	numbers[0] = 100
	fmt.Println("number in demo:", numbers[0])

}

func Test() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	var adress *int = &numbers[0] // & => adres alma operatörü

	fmt.Println(*adress) // * => adrese gitme operatörü

	pointer := unsafe.Pointer(&numbers[0])

	nextPointerInt := uintptr(pointer) + 8

	nextPointer := (*int)(unsafe.Pointer(nextPointerInt))

	fmt.Println(*nextPointer)

}
