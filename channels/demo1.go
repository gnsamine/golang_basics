package channels

func EvenNumbers(evenNumberCn chan int) {
	addition := 0
	for i := 0; i <= 10; i += 2 {

		addition = addition + i

	}
	evenNumberCn <- addition

}

func OddNumbers(oddNumberCn chan int) {
	addition := 0
	for i := 1; i <= 10; i += 2 {
		addition = addition + i

	}
	oddNumberCn <- addition

}
