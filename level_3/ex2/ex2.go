package ex2

import "fmt"

func Ex2(){
	for i:= 65; i<=90;i++{
		fmt.Println(i)
		for j:=0; j<3; j++{
			fmt.Printf("\t %#U",i)
			fmt.Println()
		}
	}
}
