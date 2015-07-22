package main

import (
	"fmt"
)

func main() {
	// Create a map of colors and color hex codes.
	colors := map[string]string{
		"AliceBlue":   "#f0f8ff",
		"Coral":       "ff7F50",
		"DarkGray":    "#a9a9a9",
		"ForestGreen": "#228b22",
	}

	// Display all the colors in the map.
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

	// removendo
	// delete(colors, "Coral")
	removeColor(colors, "Coral")

	fmt.Print("\n\n")
	for key, value := range colors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}
}

func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
