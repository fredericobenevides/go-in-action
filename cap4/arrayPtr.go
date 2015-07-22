package main

import (
	"fmt"
)

func main() {
	array := [5]*int{0: new(int), 1: new(int)}

	*array[0] = 10
	*array[1] = 20

	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])

	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}

	// Add colors
	*array2[0] = "Red"
	*array2[1] = "Blue"
	*array2[2] = "Green"

	array1 = array2
	fmt.Println(array1)
	fmt.Println(array2)
}
