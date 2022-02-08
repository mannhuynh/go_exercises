package ex3

import "fmt"

func Ex3() {
	defer foo()
	bar()
	wutang()
}

func foo() {
	fmt.Println("This is foo and will go after all because I am deferred")
}

func bar() {
	fmt.Println("I am bar and I will go before foo")
}

func wutang() {
	fmt.Println("I will go wherever I am placed!!! But before foo ")
}
