package structs

import "fmt"

func Demo1() {
	fmt.Println(product{"Bilgisayar", 15000, "SD"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
