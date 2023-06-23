package for_range

import "fmt"

func Demo3() {
	dictionary := map[string]string{"book ": "kitap", "table ": "masa"}
	for k, v := range dictionary {
		fmt.Print(k)
		fmt.Println(v)
	}

}
