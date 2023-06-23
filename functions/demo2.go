package functions

func Multiplecalculations(number1 int, number2 int) (int, int, int, float32) {

	add := number1 + number2
	substract := number1 - number2
	multiply := number1 * number2
	divide := float32(number1 / number2)

	return add, substract, multiply, divide
}
