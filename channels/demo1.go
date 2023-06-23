package channels

func EvenNumbers(EvenNumberCn chan int) {
	addition := 0
	for i := 0; i <= 10; i += 2 {

		addition = addition + 1

	}
	EvenNumberCn <- addition

}

func OddNumbers(OddNumberCn chan int) {
	addition := 0
	for i := 1; i <= 10; i += 2 {
		addition = addition + 1

	}
	OddNumberCn <- addition

}
