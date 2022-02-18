package ex1

import "fmt"

func Ex1() {
	funcLiteral()
	bufferChannal()
}

func funcLiteral() {
	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
}

func bufferChannal() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)
}
