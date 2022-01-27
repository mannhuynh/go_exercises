package ex4

import "fmt"

// Anonymous struct - required two pairs of curly brakets
func Ex4() {
	stru := struct {
		name string
		age  int
		good bool
	}{
		name: "Vuong",
		age:  30,
		good: true,
	}
	fmt.Println(stru)
}
