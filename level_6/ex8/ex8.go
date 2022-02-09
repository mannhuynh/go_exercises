package ex8

import "fmt"

func Ex8() {
	// fmt.Println("We need to call foo() and also call the variable fooFunc() to get the value form foo() function")
	fooFunc := foo()
	fmt.Println(fooFunc())
	// fmt.Println("Below functions only give the address of the function foo or address of the variables it is assigned to")
	fmt.Println(foo)
	fooFuncAddress1 := foo()
	fmt.Println(fooFuncAddress1)
	fooFuncAddress2 := foo
	fmt.Println(fooFuncAddress2)
	fooFuncAddress3 := foo
	fmt.Println(fooFuncAddress3())

}

func foo() func() int {
	return func() int {
		return 111
	}
}
