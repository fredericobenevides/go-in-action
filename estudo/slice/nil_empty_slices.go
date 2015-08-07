package main

import (
	"fmt"
)

func main() {
	var slice_1 int
	fmt.Println("\nSlice 1 - nil")
	fmt.Println(slice_1) // 0

	slice_2 := make([]int, 0)
	fmt.Println("\nSlice 2 - empty")
	fmt.Println(slice_2) // []

	slice_3 := []int{}
	fmt.Println("\nSlice 3 - empty")
	fmt.Println(slice_3) // []
}
