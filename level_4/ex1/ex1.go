package ex1

import "fmt"

func Ex1() {
	arr := [5]int{5, 2, 3, 5, 6}
	// arr[0] = 4
	// arr[1] = 7
	// arr[2] = 1
	// arr[3] = 3
	// arr[4] = 5

	for i, v := range arr {
		fmt.Println(i, v)
	}
	fmt.Printf("Type of array: %T\n", arr)
	fmt.Println("----------------")

}
