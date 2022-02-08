package ex4

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func Ex4() {

	p1 := person{
		first: "Kathy",
		last:  "Huynh",
		age:   10,
	}

	p1.speak()

}

func (p person) speak() {
	fmt.Printf("My name is %v %v and I am %d\n", p.first, p.last, p.age)
}
