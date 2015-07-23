package main

import (
	"fmt"
)

type Duration int64

func main() {
	var dur Duration

	// lança excessão, pois os tipos são incompatíveis,
	// mesmo que int64 seja a base para o Duration
	dur = int64(1000)
	fmt.Println(dur)
}
