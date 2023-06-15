package conditionals

import "fmt"

func Demo1() {

	var hesap float64 = 1000
	var çekilmekistenen float64 = 289

	if çekilmekistenen > hesap {
		fmt.Println("insufficient funds")
		fmt.Println("remainind money: " + fmt.Sprintf("%v", hesap))
		fmt.Printf("remaining money: %v ", hesap)
	}

	if çekilmekistenen <= hesap {
		fmt.Println(çekilmekistenen, "machine is giving your money")
		fmt.Println("kalan para: ", hesap-çekilmekistenen)

	}

}
func Main() {
	Demo1()
}
