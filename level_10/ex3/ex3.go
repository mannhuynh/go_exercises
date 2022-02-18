package ex3

import "fmt"

func Ex3() {
	c := gen()

	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	// Need to put the for loop into the concurency function
	go func() {
		for i := 0; i < 10; i++ {
			c <- i

		}
		// close(c) is required in here
		close(c)
	}()

	return c
}

// pull the values off the channel using a for range loop
func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
