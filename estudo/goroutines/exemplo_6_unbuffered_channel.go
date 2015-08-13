package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	var bastao = make(chan int)

	go Corredores(bastao)
	bastao <- 1

	wg.Wait()
	fmt.Println("Corrida finalizada")
}

func Corredores(bastao chan int) {
	var novoCorredor int

	corredor := <-bastao
	fmt.Printf("Corredor %d está correndo com o bastão\n", corredor)

	if corredor != 4 {
		novoCorredor = corredor + 1
		fmt.Printf("Corredor %d está na linha\n", novoCorredor)
		go Corredores(bastao)
	}

	time.Sleep(500 * time.Millisecond)

	if corredor == 4 {
		fmt.Printf("Corredor %d finalizou a corrida\n", corredor)
		wg.Done()
		return
	}

	fmt.Printf("Corredor %d pegou o bastão do %d\n\n", novoCorredor, corredor)
	bastao <- novoCorredor
}
