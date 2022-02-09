package ex7

import "fmt"

func Ex7() {
	fStr, fLen := foo("I am smart")
	fmt.Println("The string", fStr, "has the length of", fLen, "characters")
}

func foo(st string) (string, int) {
	str := st
	strLen := len(str)
	fmt.Printf("The length of the string `%v` is %d\n", str, strLen)
	return str, strLen
}
