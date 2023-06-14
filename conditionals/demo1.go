package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 1000
	var çekilmekistenen float64 = 289

	if çekilmekistenen > hesap {
		fmt.Println("hesaptaki para yetersiz")
	}
	if çekilmekistenen <= hesap {
		fmt.Println(çekilmekistenen, "veriliyor")
		fmt.Println("kalan para: ", hesap-çekilmekistenen)

	}

}
