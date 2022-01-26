package ex3

import "fmt"

func Ex3() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

	x2 := [][]int{x[:4], x[5:], x[2:7], x[2:6]}
	fmt.Println(x2)
	fmt.Println("---------------")
}
