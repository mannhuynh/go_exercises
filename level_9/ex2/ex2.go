package ex2

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) speak() {
	fmt.Println(p.name, p.age, "is speaking!!!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func Ex2() {
	//you CAN pass a value of type *person into saySomething
	p1 := person{"Vuong Huynh", 32}
	saySomething(&p1)

	// you CANNOT pass a value of type person into saySomething. speak() has pointer receiver!!!
	// saySomething(p1)

	p1.speak()

}
