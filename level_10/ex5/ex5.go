package ex5

import "fmt"

func Ex5() {
	c := make(chan int, 1)

	c <- 111

	v, ok := <-c
	fmt.Println(v, ok)

	close(c)

	v, ok = <-c
	fmt.Println(v, ok)
}
