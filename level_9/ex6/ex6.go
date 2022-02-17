package ex6

import (
	"fmt"
	"runtime"
)

func Ex6() {
	fmt.Println("go OS ", runtime.GOOS)
	fmt.Println("go ARCH ", runtime.GOARCH)

}
