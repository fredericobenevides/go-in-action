package main

import (
	"fmt"
)

func main() {
	s := make([]string, 5)
	fmt.Print("1 ")
	fmt.Println(s)

	i := make([]int, 3, 5)
	fmt.Print("2 ")
	fmt.Println(i)

	s = []string{"Red", "Blue", "Green", "Yellow", "Pink"}
	fmt.Print("3 ")
	fmt.Println(s)

	i = []int{10, 20, 30}
	fmt.Print("4 ")
	fmt.Println(i)

	// create a slice of strings.
	// Initialize the 100th element with an empty string.
	s = []string{99: ""}
	fmt.Print("5 ")
	fmt.Println(s)

	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3]
	newSlice[1] = 100
	fmt.Print("6 ")
	fmt.Println(slice)
	fmt.Println(newSlice)

	newSlice = append(newSlice, 60)
	fmt.Print("7 ")
	fmt.Println(newSlice)

	fmt.Print("\n\n\n")
	// Formulas
	// source[i:j:k] or [2:3:4]
	// length: j - i   or 3 - 2 = 1
	// capacity: k - i  or 4 -2 = 2
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	sliceSource := source[2:3:4]
	fmt.Print("8 ")
	fmt.Println(source)
	fmt.Println(sliceSource)
	sliceSource = append(sliceSource, "Teste")
	fmt.Println(source)
	fmt.Println(sliceSource)
	sliceSource = append(sliceSource, "Outro")
	fmt.Println(source)
	fmt.Println(sliceSource)

	fmt.Print("\n\n\n")
	s1 := []int{1, 2}
	s2 := []int{3, 4}
	fmt.Printf("%v\n", append(s1, s2...))
}
