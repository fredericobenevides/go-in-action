package main

import (
	"fmt"
)

func main() {
	soma := teste()
	fmt.Println(soma())
	fmt.Println(soma())
	fmt.Println(soma())
	fmt.Println(soma())
}

func teste() func() int {
	var soma int
	return func() int {
		soma += 1
		return soma
	}
}
