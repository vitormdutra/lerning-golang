package main

import (
	"fmt"
)

type v int

var x v
var y int

func main() {
	fmt.Printf("x = %v\n", x)
	fmt.Printf("x = %T\n", x)
	x = 42
	fmt.Printf("x = %v\n", x)
	y = int(x)
	fmt.Printf("y = %v\n", y)
	fmt.Printf("y = %T\n", y)
}
