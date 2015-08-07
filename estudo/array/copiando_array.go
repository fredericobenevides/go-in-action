package main

import (
	"fmt"
)

func main() {
	array_1 := [5]int{1, 2, 3, 4, 5}
	array_2 := array_1

	// Realiza uma cópia do array
	array_1[0] = 100
	fmt.Println(array_1)
	fmt.Println(array_2)
}
