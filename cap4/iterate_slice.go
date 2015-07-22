package main

import (
	"fmt"
)

func main() {
	// Create a slice of integers.
	// Contains a length and capacity of 4 elements.
	slice := []int{10, 20, 30, 40}

	// Iterate over each element and display each value.
	for index, value := range slice {
		fmt.Printf("Index: %d Value: %d\n", index, value)
	}

	// Iterate over each element and display the value and addesses.
	for index, value := range slice {
		fmt.Printf("Index: %d Value-Addr: %X ElemAddr %X\n", index, &value, &slice[index])
	}

	// Iterate over each element starting at element 3.
	for index := 2; index < len(slice); index++ {
		fmt.Printf("Index: %d Value: %d\n", index, slice[index])
	}
}
