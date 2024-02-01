package main

import (
	"fmt"
)

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB
	GB
	TB
)

func main() {

	fmt.Printf("%b\t\t\t", KB)
	fmt.Printf("%d\n", KB)
	fmt.Printf("%b\t\t\t", MB)
	fmt.Printf("%d\n", MB)
	fmt.Printf("%b\t\t\t", GB)
	fmt.Printf("%d\n", GB)
	fmt.Printf("%b\t\t\t", TB)
	fmt.Printf("%d\n", TB)

	fmt.Println()
	x := 24
	y := x >> 2
	z := x << 2
	fmt.Printf("%b\n", x)
	fmt.Printf("%b\n", y)
	fmt.Printf("%b\n", z)

}
