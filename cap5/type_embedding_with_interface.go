// Sample program to show how to embed a type into another type and
// the relationship between the inner and outer type.
package main

import (
	"fmt"
)

// notifier is an interface that defined notification
type notifier interface {
	notify()
}

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

// Disable to see only the user email
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s %s\n", a.name, a.email)
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "john smith",
			email: "john@yahoo.com",
		},
		level: "super",
	}

	// Send the admin user a notification.
	// The embedded inner type's implementation of the
	// interface is "promoted" to the outer type.
	sendNotification(&ad)

	// we can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is promoted.
	ad.notify()
}

// sendNotification accepts values that implement the notifier
// interface and sends notification
func sendNotification(n notifier) {
	n.notify()
}
