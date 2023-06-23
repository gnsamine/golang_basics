package functions

func AddVariadic(numbers ...int) (int, int) {
	// variadic functions act like as array
	addition := 0
	multiplication := 1
	for i := 0; i < len(numbers); i++ {
		addition = addition + numbers[i]

	}
	for i := 0; i < len(numbers); i++ {
		multiplication = multiplication * numbers[i]
	}
	return addition, multiplication

}
