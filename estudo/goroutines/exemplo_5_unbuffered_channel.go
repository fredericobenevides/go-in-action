package main

import (
	"fmt"
	"sync"
	// "time"
)

var wg sync.WaitGroup

func main() {
	tacada := make(chan int)

	wg.Add(2)

	go Jogador("Nadal", tacada)
	go Jogador("Guga", tacada)

	tacada <- 1

	wg.Wait()
	fmt.Println("Jogo finalizado")
}

func Jogador(nome string, tacada chan int) {
	defer wg.Done()

	bola, ok := <-tacada
	if !ok {
		fmt.Println("Bola", bola)
		return
	}
	close(tacada)

	// for {
	// 	bola, ok := <-tacada
	// 	if !ok {
	// 		return
	// 	}
	// 	if bola == 3 {
	// 		close(tacada)
	// 		return
	// 	}
	// 	tacada <- bola + 1
	// 	fmt.Println("Bola", bola)
	// }

}
