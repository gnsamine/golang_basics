package interfaces

import "fmt"

type Mortgage struct {
	CreditPaymentTotal float64
	addres             string
	rate               float64
}
type Car struct {
	CreditPaymentTotal float64
	carInfo            string
	rate               float64
}
type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.CreditPaymentTotal * m.rate / 100 / 12
}
func (c Car) Calculate() float64 {
	return c.CreditPaymentTotal * c.rate / 100 / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {

	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal = paymentTotal + credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, CreditPaymentTotal: 100000, addres: "ankara"}
	credit2 := Mortgage{rate: 12, CreditPaymentTotal: 1000000, addres: "istanbul"}
	credit3 := Car{rate: 15, CreditPaymentTotal: 50000, carInfo: "Polo"}

	credits := []CreditCalculator{credit1, credit2, credit3}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("total Payment: ", total)
}
