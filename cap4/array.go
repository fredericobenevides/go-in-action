package main

import (
	"fmt"
)

func main() {
	var array [5]int
	fmt.Print("1 ")
	fmt.Println(array)

	a := [5]int{}
	fmt.Print("2 ")
	fmt.Println(a)

	a = [5]int{10, 20, 30, 40, 50}
	fmt.Print("3 ")
	fmt.Println(a)

	a = [...]int{10, 20, 30, 40, 50}
	fmt.Print("4 ")
	fmt.Println(a)

	a = [...]int{1: 10, 4: 50}
	fmt.Print("5 ")
	fmt.Println(a)
	fmt.Print(a[1])
	fmt.Print("\n")

	a[1] = 100
	fmt.Print("6 ")
	fmt.Println(a)

	fmt.Print("\n\n")
	fmt.Println("Copiando array")
	var as [5]string
	as2 := [5]string{"Red", "Blue", "Green", "Yellow", "Pink"}

	// copy array
	as = as2
	fmt.Println(as)
	fmt.Println(as2)
	fmt.Println("\nAlterando apenas 1 array")
	as[0] = "Outro"
	fmt.Println(as)
	fmt.Println(as2)
}
