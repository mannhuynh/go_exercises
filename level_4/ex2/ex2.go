package ex2

import "fmt"

func Ex2() {
	s := []int{3, 2, 4, 1, 5, 6, 3, 4, 6, 9}

	for i, v := range s {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of Slice: %T\n", s)
	fmt.Println("---------------")
}
