package stringfunctions

import (
	"fmt"
	s "strings"
)

func Demo2() {
	name := "AmineGüneş"
	fmt.Println(s.HasPrefix(name, "Amin"))
	fmt.Println(s.HasSuffix(name, "eş"))
	fmt.Println(s.Index(name, "Gü"))
	harfler := []string{"a", "m", "i", "n"}
	fmt.Println(s.Join(harfler, ""))

	result := s.Join(harfler, "*")
	fmt.Println(result)
	fmt.Println(s.Replace(result, "*", "", 3))
	fmt.Println(s.Split(result, "*"))
	fmt.Println(s.Split(name, "-"))
	fmt.Println(s.Repeat(result, 9))
}
