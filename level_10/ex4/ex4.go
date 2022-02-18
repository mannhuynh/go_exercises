package ex4

import "fmt"

func Ex4() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		// q <- 1 to get out of the loop
		q <- 1
		close(c)
	}()

	return c
}

// pull the values off the channel using a select statement
func receive(c, q <-chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println("Value of c:", v)
		case <-q:
			return
		}
	}
}
