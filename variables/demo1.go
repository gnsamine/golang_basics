package variables

import "fmt"

func Demo1() {

	var kdv int = 10

	fmt.Println(kdv)
	var kdv2 float32 = 0.1
	fmt.Println(100 + 100*kdv2)
	kdv3 := 25

	fmt.Printf("veri tipi : %T\n", kdv3)

	var durum bool

	var metin1 string = "sd"
	var metin2 string = "amin"

	durum = metin1 == metin2

	var durum2 bool
	var metin3 string = "sd"
	var metin4 string = "amin"

	durum2 = metin3 != metin4
	fmt.Println(durum)
	fmt.Println(durum2)

	fmt.Println("Merhaba")
	fmt.Print("Merhaba")
	fmt.Print("Merhaba")
	fmt.Print("merhaba")
	fmt.Print("merhaba")
	fmt.Print("merhaba")
	fmt.Print("merhaba")
	fmt.Print("merhaba")
}
