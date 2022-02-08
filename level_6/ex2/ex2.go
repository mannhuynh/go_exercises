package ex2

import "fmt"

var si = []int{1, 2, 3, 4, 5, 6}

func Ex2() {
	fmt.Println(foo(si...))
	fmt.Println(bar(si))
}

func foo(x ...int) int {
	var sum int
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(xi []int) int {
	var sum int
	for _, v := range xi {
		sum += v
	}
	return sum
}
