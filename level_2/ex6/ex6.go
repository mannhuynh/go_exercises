package ex6

import "fmt"

func Ex6()  {
	
	const (
		a = iota
		b = iota * 365
		c = iota * 365
		d = iota * 365
	)

	const (
		x = iota
		y
		z
	)
	fmt.Println(a, b, c, d)
	fmt.Println()
	fmt.Println(x, y,z)

}
