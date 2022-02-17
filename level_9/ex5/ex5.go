package ex5

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Ex5() {
	var counter int64
	var wg sync.WaitGroup
	// var mux sync.Mutex

	wg.Add(2)

	go func() {
		for i := 0; i < 4; i++ {
			// mux.Lock()
			atomic.AddInt64(&counter, 1)
			v1 := atomic.LoadInt64(&counter)
			fmt.Println("Counter for v1:", v1)
			// mux.Unlock()
		}
		wg.Done()
	}()

	go func() {
		for j := 0; j < 5; j++ {
			// mux.Lock()
			atomic.AddInt64(&counter, 1)
			v2 := atomic.LoadInt64(&counter)
			fmt.Println("Counter for v1:", v2)
			// mux.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()

}
