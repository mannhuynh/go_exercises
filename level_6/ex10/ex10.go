package ex10

import (
	"fmt"
	"math"
)

func Ex10() {
	ca := circleArea()
	fmt.Println(ca)
}

func circleArea() float64 {
	radius := 10.
	g := func() float64 {
		// math.Pi is a float number, radius must be float also
		area := math.Pi * radius * radius
		return area
	}
	return g()
}
