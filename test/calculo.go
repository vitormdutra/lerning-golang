package main

import (
	"fmt"
)

type hotdog int

var y hotdog = 10

func main() {
	x := 10
	fmt.Printf("x: %v\n", x)

	x = int(y)
	fmt.Printf("b: %v", x)
}
