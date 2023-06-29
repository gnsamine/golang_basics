package interfaces

import (
	"fmt"
	"math"
)

type shape interface {
	//write methodes into interfaces
	area() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.height * r.width

}
func school(s shape) {
	fmt.Println(s.area())
}
func Demo1() {
	r := rectangle{width: 10, height: 6}//rectangleın bellekteki yerini shape tutyor
	school(r)
	k := circle{radius: 5}
	school(k)

}
