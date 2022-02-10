package ex2

import "fmt"

type person struct {
	first string
	last  string
}

func changeMe(p *person) {
	*&p.first = "Kathy"
	// *&x will simplify to x. It will not copy x
	p.last = "Huynh"
}

func Ex2() {

	p1 := person{
		first: "Cindy",
		last:  "Mai",
	}
	fmt.Println("p1 befor:", p1.first, p1.last)

	changeMe(&p1)

	fmt.Println("p1 after:", p1.first, p1.last)

}
