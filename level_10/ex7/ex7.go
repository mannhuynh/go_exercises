package ex7

import (
	"fmt"
	"sync"
)

func Ex7() {
	wg := sync.WaitGroup{}
	c := make(chan int)

	x := 10
	go func() {
		for i := 0; i < x; i++ {
			wg.Add(1)
			go func() {
				for j := 0; j < x; j++ {
					c <- j
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()

	for k := range c {
		fmt.Println("value of c:", k)
	}
}
