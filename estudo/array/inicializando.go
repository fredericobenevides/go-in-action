package main

import (
	"fmt"
)

func main() {
	fmt.Println("\nArray 1")
	var array [5]int
	fmt.Println(array) // [0, 0, 0, 0, 0]

	fmt.Println("\nArray 2")
	array_2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array_2)      // [1, 2, 3, 4, 5]
	fmt.Println(len(array_2)) // 5
	fmt.Println(cap(array_2)) // 5

	fmt.Println("\nArray 3")
	array_3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array_3)      // [1, 2, 3, 4, 5]
	fmt.Println(len(array_3)) // 5
	fmt.Println(cap(array_3)) // 5

	// Realiza uma cópia
	array_4 := array_3
	fmt.Print("\nArray 4\n")
	fmt.Println(array_4)      // [1, 2, 3, 4, 5]
	fmt.Println(len(array_4)) // 5
	fmt.Println(cap(array_4)) // 5

	array_5 := [5]int{1: 200, 4: 500}
	fmt.Println(array_5)      // [1, 200, 3, 4, 500]
	fmt.Println(len(array_5)) // 5
	fmt.Println(cap(array_5)) // 5
}
