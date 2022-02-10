package ex3

import "fmt"

// This is the bonus code for understanding pass by value and pass by pointer to a function

func Ex3() {
	x := 0
	fmt.Println("x befor:", x)
	foo(x)
	// x is unchanged after foo is called
	fmt.Println("x after:", x)

	fmt.Println("---")

	y := 0
	fmt.Println("y befor:", y)
	bar(&y)
	// y is changed after bar is called because it is passed by the pointer
	fmt.Println("y after:", y)

}

// w is passed as a non-pointer type integer
func foo(w int) {
	fmt.Println("w befor:", w)
	w = 111
	fmt.Println("w after:", w)
}

// z is passed as pointer type integer
func bar(z *int) {
	fmt.Println("z befor:", *z)
	*z = 222
	fmt.Println("z after:", *z)
}
