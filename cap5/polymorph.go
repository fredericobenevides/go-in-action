package main

import (
	"fmt"
)

// notifier is an interface that defined notification
// type behaviour
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
	name  string
	email string
}

// notify implements the notifier interface with a pointer receiver.
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s %s\n", a.name, a.email)
}

func main() {
	// Create a value of type User and send a notification
	bill := user{"Bill", "bill@email.com"}
	sendNotification(&bill)

	// Create an admin value and pass it to sendNotification
	lisa := admin{"Lisa", "lisa@email.com"}
	sendNotification(&lisa)
}

// sendNotification accepts values that implement the notifier
// interface and sends notification
func sendNotification(n notifier) {
	n.notify()
}
