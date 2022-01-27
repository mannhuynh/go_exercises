package ex3

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheels bool
}

type sedan struct {
	vehicle
	luxury bool
}

func Ex3() {
	tr := truck{
		fourWheels: true,
		vehicle: vehicle{doors: 4,
			color: "Black",
		},
	}

	se := sedan{
		vehicle: vehicle{doors: 4, color: "green"},
		luxury:  true,
	}

	fmt.Println(tr)
	fmt.Printf("Doors: %v, color: %v, four wheels? %v", tr.doors, tr.color, tr.fourWheels)
	fmt.Println()

	fmt.Println(se)
	fmt.Printf("Doors: %v, color: %v, luxury? %v", se.doors, se.color, se.luxury)
	fmt.Println()

}
