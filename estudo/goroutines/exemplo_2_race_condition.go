package main

import (
	"fmt"
	"sync"
	"time"
)

var contador = 0
var wg sync.WaitGroup

func inc(name string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		valor := contador
		time.Sleep(1000 * time.Millisecond)

		valor++
		contador = valor

		fmt.Printf("%s incrementou o valor para %d\n", name, contador)

	}
}

func main() {
	wg.Add(2)

	go inc("G1")
	go inc("G2")

	wg.Wait()
	fmt.Printf("\n\nO contador atual terminou com %d", contador)
}
