package main

import (
	"fmt"
)

func main() {
	var map_1 map[string]string
	fmt.Println("\nMap 1 - empty")
	fmt.Println(map_1)

	map_2 := make(map[string]string)
	fmt.Println("\nMap 2 - empty")
	fmt.Println(map_2)

	map_3 := map[string]string{}
	fmt.Println("\nMap 3 - empty")
	fmt.Println(map_3)
}
