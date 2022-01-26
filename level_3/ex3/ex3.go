package ex3

import (
	"fmt"
	"time"
)

func Ex3(){
	birth_year:=1980
	this_year,_,_:= time.Now().Date()
	for birth_year <= this_year{
		fmt.Println(birth_year)
		birth_year++
	}
	fmt.Println("-----------------")
}
