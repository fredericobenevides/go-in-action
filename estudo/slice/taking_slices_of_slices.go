package main

import (
	"fmt"
)

func main() {

	slice_1 := []int{1, 2, 3, 4, 5}
	slice_2 := slice_1[0:3]

	fmt.Println("Slice 1")
	fmt.Println(slice_1) // [1, 2, 3, 4, 5]

	fmt.Println("Slice 2")
	fmt.Println(slice_2) // [1, 2, 3]

	fmt.Println("Alterando o slice 1")
	slice_1[0] = 100
	slice_1[1] = 200
	slice_1[2] = 300

	fmt.Println("Imprimindo os slices depois de alterar o slice 1")

	fmt.Println("Slice 1")
	fmt.Println(slice_1) // [100, 200, 300, 4, 5]

	fmt.Println("Slice 2")
	fmt.Println(slice_2) // [100, 200, 300]

	fmt.Println("Tamanho do slice 1")
	fmt.Println(len(slice_1)) // 5
	fmt.Println(cap(slice_1)) // 5

	fmt.Println("Tamanho do slice 3")
	slice_3 := slice_1[0:3:4]
	fmt.Println(len(slice_3)) // 3
	fmt.Println(cap(slice_3)) // 4

	fmt.Println("Alterando um slice que tem uma capacidade maior que o seu tamanho")
	slice_4 := slice_1[0:3]
	fmt.Println("Adicionado um valor no slice")
	slice_4 = append(slice_4, 888)
	fmt.Println(slice_1)
	fmt.Println(slice_4)
	fmt.Println("Bug no slice1, pois utilizou a mesma referencia dele ao dar append")

	fmt.Println("Por isso deve sempre especificar a capacidade do slice")
	slice_5 := slice_1[0:3:3]
	slice_5 = append(slice_5, 999)
	fmt.Println(slice_1)
	fmt.Println(slice_5)
	fmt.Println("Funcionou sem problemas, foi alterado o slice2, mas o valor do slice1 se manteve")
}
