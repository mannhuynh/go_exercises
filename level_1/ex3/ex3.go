package ex3

import "fmt"

var x = 42
var y = "Golang"
var z = true

func Ex3(){
	s := fmt.Sprint(x, y, z)
	fmt.Println("This is s", s)
	n, err := fmt.Println(s)
	fmt.Println("fmt.Println return an int and an err", n, err)
	m,_ := fmt.Println("Try if it prints twice??!!!")
	fmt.Println(m)
}
