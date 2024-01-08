package main

import (
	"fmt"
	"math"
)

var a, b, c float64

func main() {
	fmt.Printf("Start calculator\n")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	delta := math.Pow(b, 2) - (4 * a * c)

	if delta < 0 {
		fmt.Println("no real roots")
		return
	}

	x1 := (-b + math.Sqrt(delta)) / (2 * a)
	x2 := (-b - math.Sqrt(delta)) / (2 * a)

	fmt.Printf("%v\n", x1)
	fmt.Printf("%v\n", x2)
}
