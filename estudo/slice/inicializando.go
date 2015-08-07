package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slice 1")
	slice_1 := make([]int, 5)
	fmt.Println(slice_1)      // [0, 0, 0, 0, 0]
	fmt.Println(len(slice_1)) // 5
	fmt.Println(cap(slice_1)) // 5

	fmt.Println("\nSlice 2")
	slice_2 := make([]int, 3, 5)
	fmt.Println(slice_2)      // [0, 0, 0]
	fmt.Println(len(slice_2)) // 3
	fmt.Println(cap(slice_2)) // 5

	fmt.Println("\nSlice 3")
	slice_3 := []int{1, 2, 3}
	fmt.Println(slice_3)      // [1, 2, 3]
	fmt.Println(len(slice_3)) // 3
	fmt.Println(cap(slice_3)) // 3

	fmt.Println("\nSlice 4")
	slice_4 := []int{0: 100, 2, 300}
	fmt.Println(slice_4)      // [100, 2, 300]
	fmt.Println(len(slice_4)) // 3
	fmt.Println(cap(slice_4)) // 3
}
