package ex1

import "fmt"

func Ex1() {
	fmt.Println(foo())
	fmt.Println(bar())

}

func foo() int {
	return 4
}

func bar() (int, string) {
	return 3, "Hello"
}
