package main

import (
	"fmt"
)

func main() {
	// var array [4][2]int
	array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

	fmt.Println(array)

	array = [4][2]int{1: {0: 20}, 3: {1: 41}}
	fmt.Println(array)

	var array1 [2][2]int
	var array2 [2][2]int

	array2[0][0] = 10
	array2[0][1] = 20
	array2[1][0] = 30
	array2[1][1] = 40

	array1 = array2
	fmt.Println(array1)

	var array3 [2]int = array1[1]
	fmt.Println(array3)
}
