package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func save(p product) {
	fmt.Println("product saved:", p.productName)
}
