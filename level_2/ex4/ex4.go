package ex4

import "fmt"

func Ex4(){
	a := 31
	fmt.Printf("%d\t %b\t %#x",a,a,a)
	fmt.Println()

	b := a << 1
	fmt.Printf("%d\t %b\t %#x",b,b,b)
	fmt.Println()

}
