package ex1

import "fmt"

type person struct {
	first_name  string
	last_name   string
	fav_flavors []string
}

func Ex1() {
	p1 := person{
		first_name:  "Vuong",
		last_name:   "Huynh",
		fav_flavors: []string{"Vanilla", "Coffee", "Chocolate"},
	}
	p2 := person{
		first_name:  "Kathy",
		last_name:   "Huynh",
		fav_flavors: []string{"Vanilla", "Milk Coffee", "Caramel"},
	}

	fmt.Println("Name:", p1.first_name, p1.last_name)
	fmt.Println("Favorite flavors are:")
	for _, v := range p1.fav_flavors {
		fmt.Println("\t -", v)
	}

	fmt.Println()

	fmt.Println("Name:", p2.first_name, p2.last_name)
	fmt.Println("Favorite flavors are:")
	for _, v := range p2.fav_flavors {
		fmt.Println("\t -", v)
	}

}
