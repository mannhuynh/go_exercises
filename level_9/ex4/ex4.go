package ex4

import (
	"fmt"
	"sync"
)

func Ex4() {
	var counter int
	var wg sync.WaitGroup
	var mux sync.Mutex

	wg.Add(2)

	go func() {
		for i := 0; i < 3; i++ {
			mux.Lock()
			v1 := counter
			v1++
			counter = v1
			fmt.Println("Counter for v1:", counter)
			mux.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for j := 0; j < 4; j++ {
			mux.Lock()
			v2 := counter
			v2++
			counter = v2
			fmt.Println("Counter for v2:", counter)
			mux.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()

}
