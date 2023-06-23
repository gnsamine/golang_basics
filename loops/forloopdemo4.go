package loops

import "fmt"

func Demo4() {
	var givennumber int

	fmt.Println("Write any positive number")
	fmt.Scanln(&givennumber)

	Prime := true
	for devisive := 2; devisive < givennumber; devisive++ {
		if givennumber%devisive == 0 {
			Prime = false
			break
		}

	}
	if Prime == true {
		fmt.Println("Your number is a prime number")

	} else {
		fmt.Println("your number is not a prime number")
	}
}
