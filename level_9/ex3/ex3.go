package ex3

import (
	"fmt"
	"runtime"
	"sync"
)

func Ex3() {
	var counter int
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for i := 0; i < 5; i++ {
			v1 := counter
			runtime.Gosched()
			v1++
			counter = v1
			fmt.Println("Counter for v1:", counter)
		}
		wg.Done()
	}()

	go func() {
		for j := 0; j < 10; j++ {
			v2 := counter
			runtime.Gosched()

			v2++
			counter = v2
			fmt.Println("Counter for v2:", counter)
		}
		wg.Done()
	}()
	wg.Wait()

}
