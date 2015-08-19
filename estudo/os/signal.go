package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	// signal.Stop(c)

	for {
		fmt.Println("loop")
		select {
		case <-c:
			fmt.Println("Saindo")
			return
		default:
			fmt.Println("Continuando")
		}

	}

	// // Block until  a signal is received
	// s := <-c

	// s = <-c
	// fmt.Println("Got signal:", s)
}
