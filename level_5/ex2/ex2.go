package ex2

import "fmt"

type person struct {
	first_name  string
	last_name   string
	fav_flavors []string
}

func Ex2() {
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

	m := map[string]person{
		"vuong": p1,
		"kathy": p2,
	}

	for _, v := range m {
		fmt.Println("Name: ", v.first_name, v.last_name)
		fmt.Println("Favorite flavors:")
		for _, v2 := range v.fav_flavors {
			fmt.Println("\t-", v2)
		}
		fmt.Println()
	}
}
