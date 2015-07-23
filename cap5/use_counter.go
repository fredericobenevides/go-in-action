// Sample program to show how the program can't access an
// unexported identifier from another package.
package main

import (
	"fmt"
	"go-in-action/cap5/counters"
)

func main() {
	// Create a variable of the unexported type and initialize
	// the value to 10.
	// counter := counters.alertCounter(10)

	// Create a variable of the unexported type using the exported
	// New function from the package counters.
	counter := counters.New(10)

	fmt.Printf("Counter %d\n", counter)
}
