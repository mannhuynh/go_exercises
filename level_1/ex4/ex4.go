package ex4

import "fmt"

type new_int int

var x new_int

func Ex4() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
}
