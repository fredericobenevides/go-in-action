package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var contador int32
var wg sync.WaitGroup

func inc(name string) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		atomic.AddInt32(&contador, 1)
		time.Sleep(1000 * time.Millisecond)

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
