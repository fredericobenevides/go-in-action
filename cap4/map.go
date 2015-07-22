package main

import (
	"fmt"
)

func main() {
	dict1 := make(map[string]int)
	fmt.Println(dict1)

	dict2 := map[string]string{"Red": "da1337", "Orange": "#e95a22"}
	fmt.Println(dict2)

	// não funciona
	// dict3 := map[[]string]int{}
	// fmt.Println(dict3)

	dict4 := map[int][]string{}
	fmt.Println(dict4)

	// Create a nil map by just declaring the map.
	var colors map[string]string

	// Add the Red color code to the map.
	// Não funciona pois não existe a key Red
	// colors["Red"] = "#da1337"

	// Retrieve the value for the key "Blue"
	value, exists := colors["Blue"]

	// Did the key exist?
	if exists {
		fmt.Println(value)
	}

	// ou usando assim
	value := colors["Blue"]

	// Did the key exist?
	if value != "" {
		fmt.Println(value)
	}
}
