package main

import (
	"fmt"
)

func main() {

	fmt.Println("Map 1")
	map_1 := make(map[string]string)
	fmt.Println(map_1)

	fmt.Println("Map 2")
	map_2 := map[string]string{"Red": "Vermelho", "Black": "Preto"}
	fmt.Println(map_2)

	fmt.Println("Existe a chave Red?")
	valor, existe := map_2["Red"]
	if existe {
		fmt.Println("Existe a chave \"Red\"")
		fmt.Println(valor)
	}
}
