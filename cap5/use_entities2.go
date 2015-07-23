// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly
package main

import (
	"fmt"
	"go-in-action/cap5/entities2"
)

func main() {
	a := entities.Admin{
		Rights: 10,
	}

	// Set the exported fields from the unexported
	// inner type.
	a.Name = "Bill"
	a.Email = "bill@email.com"

	fmt.Printf("User: %v\n", a)

	// Não funciona, já que user não foi exportado
	// a.user.Name = "Bill"
	// a.user.Email = "bill@email.com"
}
