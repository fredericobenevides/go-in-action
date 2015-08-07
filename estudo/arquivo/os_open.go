package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Lendo arquivo")

	file, err := os.Open("texto.txt")
	if err != nil {
		log.Fatalln("Não foi possível ler o arquivo")
	}

	fmt.Println("\n\nUsando o for para ler os arquivos")

	for err == nil {
		buffer := make([]byte, 10)

		_, err = file.Read(buffer)
		fmt.Print(string(buffer))
	}
}
