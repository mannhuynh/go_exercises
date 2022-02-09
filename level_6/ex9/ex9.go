package ex9

import "fmt"

func Ex9() {
	// Do not call the callback function, intead just past it in
	foo(iamCallback)
	Ex9a()
}

func iamCallback() int {
	return 111
}

func foo(f func() int) {
	// Remember to call the function f() in the print statement
	fmt.Println("The value of callback function is:", f())
}
