package ex5

import "fmt"

func Ex5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)

	y := append(x[:3], x[6:]...)
	fmt.Println(y)
	fmt.Println("---------------")

}
