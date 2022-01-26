package ex9

import "fmt"

func Ex9(){
	favSport := "soccer"
	switch favSport{
	case "football": fmt.Println("football is my favorite sport")
	case "tennis": fmt.Println("I don't know tennis")
	case "soccer": fmt.Println("I love soccer")
	default: fmt.Println("I love all kind of sports")
	}
	fmt.Println("-------------")
}
