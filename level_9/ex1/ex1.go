package ex1

import (
	"fmt"
	"runtime"
	"sync"
)

func Ex1() {
	fmt.Println("First check CPU:\t", runtime.NumCPU())
	fmt.Println("First check GoRoutine:\t", runtime.NumGoroutine())

	// wg := sync.WaitGroup{}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("FIRST additional goroutine")
		wg.Done()
	}()

	go func() {
		fmt.Println("SECOND additional goroutine")
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Last check CPU:\t", runtime.NumCPU())
	fmt.Println("Last check GoRoutine:\t", runtime.NumGoroutine())

}
