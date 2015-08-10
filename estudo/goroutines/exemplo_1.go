package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Iniciando goroutine")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("Esperando terminar")
	wg.Wait()

	fmt.Println("Terminando")
}
