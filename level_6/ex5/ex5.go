package ex5

import (
	"fmt"
	"math"
)

// create a type SQUARE
type square struct {
	len float64
}

// create a type CIRCLE
type circle struct {
	radius float64
}

// attach a method to each that calculates AREA and returns it
func (s square) area() float64 {
	return s.len * s.len
}

// attach a method to each that calculates AREA and returns it
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// create a type SHAPE that defines an interface as anything that has the AREA method
// Need to put the float64 to area() method.
type shape interface {
	area() float64
}

// create a func INFO which takes type shape and then prints the area
func info(sh shape) {
	switch sh.(type) {
	case square:
		fmt.Println("The area of square is:", sh.(square).area())
	case circle:
		fmt.Println("The area of circle is:", sh.(circle).area())
	}
}

// func info(sh shape) {
// 	fmt.Println(sh.area())
// }

func Ex5() {
	s := square{len: 5.0}
	c := circle{radius: 2.0}

	info(s)
	info(c)
}
