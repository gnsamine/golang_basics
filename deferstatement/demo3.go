package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func Log() {
	fmt.Println("logged")
}

func (p product) save() {
	fmt.Println("product is saved:", p.productName, p.unitPrice)
	defer Log() //we should put defer in our code.
	//Even if  other codes dont work, log function will work when defer is used
}

func Demo3() {
	p := product{productName: "laptop", unitPrice: 50000}
	defer p.save()
	p = product{productName: "mouse", unitPrice: 50}

	fmt.Println("succesfull")
	fmt.Println(p.productName)
	//defer gives us what is written before defer
	//p.product name gives us "mouse" because its value is mouse now.
	// if we write defer in 16th line we get laptop,
	// if we write defer in 18th line we get mouse

}
