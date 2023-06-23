package structs

import "fmt"

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() { // this is method. method can be differentiated from function by ( c customer) part
	fmt.Println("added:", c.firstName)

}
func (c customer) update() {
	fmt.Println("updated:", c.lastName)
}

func Demo2() {
	c := customer{firstName: "sd", lastName: "d", age: 35}
	c.save()
	c.update()
}
