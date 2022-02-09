package ex6

import "fmt"

func Ex6() {
	func(st string) (string, int) {
		str := st
		strLen := len(str)
		fmt.Printf("The length of the string `%v` is %d\n", str, strLen)
		return str, strLen
	}("I am smart")
}
