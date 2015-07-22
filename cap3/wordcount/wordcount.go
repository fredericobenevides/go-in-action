package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"go-in-action/cap3/words"
)

func main() {
	filename := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text. \n", count)
}
