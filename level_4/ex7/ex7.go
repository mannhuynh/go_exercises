package ex7

import "fmt"

func Ex7() {
	x2 := make([][]string, 2)

	x2[0] = []string{"James", "Bond", "Shaken, not stirred"}
	x2[1] = []string{"Miss", "Moneypenny", "Helloooooo, James."}

	fmt.Println(x2)
	fmt.Println("------------------")
}
