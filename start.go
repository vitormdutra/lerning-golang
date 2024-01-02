package main

import (
	"fmt"
)

func main() {
	x := 10
	y := "test"
	fmt.Printf("x: %v, %T,\n", x, x)
	fmt.Printf("y: %v, %T,\n\n\n", y, y)
	numerodebytes, erros := fmt.Print("Hello World!")
	fmt.Println(numerodebytes, erros)

	x = 20
	fmt.Printf("x: %v, %T,\n", x, x)

	t := 10
	qualquercoisa(t)
}

func qualquercoisa(x int) {
	fmt.Println(x)
}
