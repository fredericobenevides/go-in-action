package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	contents, err := ioutil.ReadFile("texto.txt")
	if err != nil {
		fmt.Println("Não foi possível ler o arquivo")
	}

	fmt.Printf("Conteúdo do arquivo:\n%s", string(contents))
}
