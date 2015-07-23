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
// se colocar ponteiro, dá erro, pois para ser pointeiro
// o objeto também "user" deve ser "&user"
// func (u *user) notify() {
func (u user) notify() {
	fmt.Printf("Sending user email to %s %s\n", u.name, u.email)
}

func main() {
	// Create a value of type User and send a notification
	u := user{"Bill", "bill@email.com"}

	// funciona tanto o método notify receber (u user)
	// ou (u *user)
	sendNotification(u)
	// sendNotification(&u)
}

// sendNotification accepts values that implement the notifier
// interface and sends notification
func sendNotification(n notifier) {
	n.notify()
}
