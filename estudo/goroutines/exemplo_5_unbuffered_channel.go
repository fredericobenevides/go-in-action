package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tacada := make(chan int)

	wg.Add(2)

	// go Jogador1("Nadal", tacada)
	// go Jogador2("Guga", tacada)

	// go Jogador("Nadal", tacada)
	// go Jogador("Guga", tacada)

	go JogadorSemChamarClose("Nadal", tacada)
	go JogadorSemChamarClose("Guga", tacada)

	// necessário para executar os goroutines
	tacada <- 1

	wg.Wait()
	fmt.Println("\nJogo finalizado")
}

func Jogador(name string, tacada chan int) {
	defer wg.Done()

	for {
		bola, ok := <-tacada
		if !ok {
			fmt.Printf("Jogador %s ganhou o jogo!!!!!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%10 == 0 {
			fmt.Printf("*** Jogador %s errou a bola *** \n", name)
			close(tacada)
			return
		} else {
			fmt.Printf("Jogador %s bateu na bola %d\n", name, bola)
		}

		bola++
		tacada <- bola
	}
}

// Nessa função, não é chamado close, porém como é fechado o WaitGroup
// nessa própria função quando o jogador erra a bola,
// a outra goroutine não é executada para avisar que o jogador ganhou o jogo
func JogadorSemChamarClose(name string, tacada chan int) {
	for {
		bola, ok := <-tacada
		if !ok {
			fmt.Printf("Jogador %s ganhou o jogo!!!!!\n", name)
			return
		}

		n := rand.Intn(100)
		if n%10 == 0 {
			fmt.Printf("*** Jogador %s errou a bola *** \n", name)
			// tem que fazer duas vezes para matar
			// a goroutine corrente e a outra que está aguardando retorno
			wg.Done()
			wg.Done()
			return
		} else {
			fmt.Printf("Jogador %s bateu na bola %d\n", name, bola)
			bola++
			tacada <- bola
		}
	}
}

func Jogador1(nome string, tacada chan int) {
	defer wg.Done()

	fmt.Println("Jogador 1", nome)
	bola := <-tacada
	fmt.Println("Tacada", bola)

	close(tacada)
}

func Jogador2(nome string, tacada chan int) {
	defer wg.Done()

	fmt.Println("Jogador 2", nome)
	bola := <-tacada
	fmt.Println("Tacada", bola)
}

// func Jogador(nome string, tacada chan int) {
// 	defer wg.Done()

// 	bola, ok := <-tacada
// 	if !ok {
// 		fmt.Println("Bola", bola)
// 		return
// 	}
// 	close(tacada)

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

// }
