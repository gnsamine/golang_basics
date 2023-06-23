package loops

import "fmt"

//finding if given number is prime number or not.

func Demo3() {
	givennumber := 0
	fmt.Println("put any positive number that 'x>=2'")
	fmt.Scanln(&givennumber)

	prime := true
	for devisive := 2; devisive < givennumber; devisive++ {
		if givennumber%devisive == 0 {
			prime = false
		}

	}

	if prime == true {
		fmt.Println("your number is Prime")
	} else {
		fmt.Println("your number is not prime")
	}

}
