package ex4

import (
	"fmt"
	"time"
)

func Ex4(){
birth_year := 1980
this_year, _, _ := time.Now().Date()

for {
	if birth_year > this_year{
		break
	}
	fmt.Println(birth_year)
	birth_year++
}
fmt.Println("-----------")
}
