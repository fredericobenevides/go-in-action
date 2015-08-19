package main

import (
	"fmt"
	"sync"
	"time"
)

type testeType struct {
	valor int
}

var wg sync.WaitGroup

func main() {
	fmt.Println("Testando channel")
	// wg.Add(1)

	teste := make(chan testeType)
	go testeChan(teste)
	// teste <- testeType{valor: 1}

	// wg.Wait()

	for {
		fmt.Println("loop")

		select {
		case <-teste:
			fmt.Println("Saindo do loop")
			return
		default:
			fmt.Println("Continuando no loop")
		}
	}
}

func testeChan(t chan testeType) {
	// defer wg.Done()
	fmt.Println("Esperando por alguns segundos")
	time.Sleep(2 * time.Second)
	fmt.Println("Finalizando a espera")
	t <- testeType{valor: 1}
	fmt.Println("Aguardando receber t")
	valor := <-t
	fmt.Println(valor)
	fmt.Println(valor.valor)
}
