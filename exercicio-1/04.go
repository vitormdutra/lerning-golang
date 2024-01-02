package main

import (
	"fmt"
)

type v int

var x v

func main() {
	fmt.Printf("x = %v\n", x)
	fmt.Printf("x = %T\n", x)
	x = 42
	fmt.Printf("x = %v\n", x)
	h, err := fmt.Print(x)
	fmt.Println(h, err)
}
