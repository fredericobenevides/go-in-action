package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	for i := 0; i < 10; i++ {
		go teste(channel)
		channel <- i
	}
	for i := 0; i < 10; i++ {
		// channel <- i
		// wg.Done()
	}

	fmt.Println("Finalizando")
}

func teste(t chan int) {
	valor := <-t
	fmt.Println(valor)
}
