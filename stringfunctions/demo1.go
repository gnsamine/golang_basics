package stringfunctions

//alis
import (
	"fmt"
	s "strings"
)

//case sensitive: Big and small case aare different
//ascii

func Demo1() {
	name := "amine"
	//how many a is in the world amine
	fmt.Println(s.Count(name, "a"))
	//does amine contains k or a
	fmt.Println(s.Contains(name, "k"))
	fmt.Println(s.Contains(name, "a"))
	//Index returns the index of the first instance of
	//substr in s, or -1 if substr is not present in s.
	fmt.Println(s.Index(name, "e"))
	fmt.Println(s.Index(name, "k"))
	result := s.Index(name, "k")
	if result != -1 {
		fmt.Println("it is in the word")
	} else {
		fmt.Println("it is not in the word")
	}
	fmt.Println(s.ToLower(name))
	fmt.Println(s.ToUpper(name))
}
