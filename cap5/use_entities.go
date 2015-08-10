// Sample program to show how unexported fields from an exported
// struct type can't be accessed directly
package main

import (
	"fmt"
	"go-in-action/cap5/entities"
)

func main() {
	u := entities.User{
		Name:  "Bill",
		email: "bill@email.com", //email minúsculo não é exportado
	}

	fmt.Printf("User: %v\n")
}
