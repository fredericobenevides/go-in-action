package main

import (
	"fmt"
)

func main() {
	array := [...]*int{0: new(int), 1: new(int)}

	*array[0] = 10
	*array[1] = 20

	fmt.Println("Imprimindo o valor do array sem acessar o ponteiro")
	fmt.Println(array[0])
	fmt.Println(array[1])

	fmt.Println("Imprimindo o valor do array acessando o ponteiro")
	fmt.Println(*array[0])
	fmt.Println(*array[1])
}
