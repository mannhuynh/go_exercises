package ex10

import "fmt"

func Ex10() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	m["vuong_huynh"] = []string{"Kathy", "Computer", "Hate Kathy the most"}

	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k, v)
	}

}
