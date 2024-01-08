package main

import (
	"fmt"
	"math"
)

var a, b, c, d float64

func main() {
	fmt.Println("start")
	a := 1.0
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&d)

	s1 := b + c + d
	fmt.Printf("s1 = %v\n", s1)
	s2 := b*c + b*d + c*d
	fmt.Printf("s2 = %v\n", s2)
	p := b * c * d
	fmt.Printf("p = %v\n", p)

	result := math.Pow(a, 3) - math.Pow(s1, 2) + s2 - p
	fmt.Printf("result = %v", result)
}
