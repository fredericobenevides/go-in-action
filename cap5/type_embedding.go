// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import (
	"fmt"
)

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a pointer receiver.
func (u *user) notify() {
	fmt.Printf("Sending user email to %s %s\n", u.name, u.email)
}

// admin defines a admin in the program.
type admin struct {
	user  // Embedded Type
	level string
}

// func (a *admin) notify() {
// 	fmt.Printf("Sending admin email to %s %s\n", a.name, a.email)
// }

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// we can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}
