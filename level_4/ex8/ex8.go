package ex8

import "fmt"

func Ex8() {
	m := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	fmt.Println(m)
	for k, v := range m {
		fmt.Println("key", k, "has value of:")
		for i, v := range v {
			fmt.Println("\t", i, v)
		}

	}
	fmt.Println("-----------")
}
